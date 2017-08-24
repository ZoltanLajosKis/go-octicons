package octicons

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
)

func TestToSVGNil(t *testing.T) {
	for _, name := range Symbols() {
		func(name string) {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				octi := Octicons(name)
				goXML := octi.ToSVG(nil)
				nodeXML := nodeCmd(t, "require('octicons')[\""+name+"\"].toSVG()")
				compareXML(t, goXML, nodeXML)
			})
		}(name)
	}
}

func TestToSVGUseNil(t *testing.T) {
	for _, name := range Symbols() {
		func(name string) {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				octi := Octicons(name)
				goXML := octi.ToSVGUse(nil)
				nodeXML := nodeCmd(t, "require('octicons')[\""+name+"\"].toSVGUse()")
				compareXML(t, goXML, nodeXML)
			})
		}(name)
	}
}

func TestOptions(t *testing.T) {
	compareXML(t,
		Alert.ToSVG(Opts{"class": "close"}),
		nodeCmd(t, "require('octicons').alert.toSVG({\"class\": \"close\"})"))

	compareXML(t,
		ArrowDown.ToSVGUse(Opts{"class": "close"}),
		nodeCmd(t, "require('octicons')[\"arrow-down\"].toSVGUse({\"class\": \"close\"})"))

	compareXML(t,
		ArrowLeft.ToSVG(Opts{"aria-label": "Close the window"}),
		nodeCmd(t, "require('octicons')[\"arrow-left\"].toSVG({\"aria-label\": \"Close the window\"})"))

	compareXML(t,
		ArrowRight.ToSVGUse(Opts{"aria-label": "Close the window"}),
		nodeCmd(t, "require('octicons')[\"arrow-right\"].toSVGUse({\"aria-label\": \"Close the window\"})"))

	for n := 0; n <= 100; n++ {
		func(n int) {
			t.Run(strconv.Itoa(n), func(t *testing.T) {
				t.Parallel()

				compareXML(t,
					ArrowUp.ToSVG(Opts{"width": strconv.Itoa(n)}),
					nodeCmd(t, "require('octicons')[\"arrow-up\"].toSVG({\"width\": \""+strconv.Itoa(n)+"\"})"))

				compareXML(t,
					Broadcast.ToSVGUse(Opts{"width": strconv.Itoa(n)}),
					nodeCmd(t, "require('octicons').broadcast.toSVGUse({\"width\": \""+strconv.Itoa(n)+"\"})"))

				compareXML(t,
					Calendar.ToSVG(Opts{"height": strconv.Itoa(n)}),
					nodeCmd(t, "require('octicons').calendar.toSVG({\"height\": \""+strconv.Itoa(n)+"\"})"))

				compareXML(t,
					Clippy.ToSVGUse(Opts{"height": strconv.Itoa(n)}),
					nodeCmd(t, "require('octicons').clippy.toSVGUse({\"height\": \""+strconv.Itoa(n)+"\"})"))
			})
		}(n)
	}

	compareXML(t,
		Dash.ToSVG(Opts{"width": "45", "height": "60"}),
		nodeCmd(t, "require('octicons').dash.toSVG({\"width\": \"45\", \"height\": \"60\"})"))

	compareXML(t,
		Database.ToSVGUse(Opts{"width": "45", "height": "60"}),
		nodeCmd(t, "require('octicons').database.toSVGUse({\"width\": \"45\", \"height\": \"60\"})"))

	compareXML(t,
		Diff.ToSVGUse(Opts{"width": "not_a_number"}),
		nodeCmd(t, "require('octicons').diff.toSVGUse({})"))

	compareXML(t,
		Eye.ToSVG(Opts{"height": "not_a_number"}),
		nodeCmd(t, "require('octicons').eye.toSVG({})"))
}

func TestInterface(t *testing.T) {
	alert, ok := Alert.(*octicon)
	assertEqual(t, ok, true)

	assertEqual(t, Alert.Symbol(), alert.symbol)
	assertEqual(t, Alert.Keywords(), alert.keywords)
	assertEqual(t, Alert.Path(), alert.path)
	assertEqual(t, Alert.Options(), alert.options)
	assertEqual(t, Alert.Width(), alert.width)
	assertEqual(t, Alert.Height(), alert.height)
}

func TestOcticons(t *testing.T) {
	assertEqual(t, Octicons("logo-github"), LogoGithub)
	assertEqual(t, Octicons("no-such-octicon-symbol"), nil)
}

func TestCached(t *testing.T) {
	name := "logo-github"
	octi := Octicons(name)

	goXML := octi.ToSVG(nil)
	nodeXML := nodeCmd(t, "require('octicons')[\""+name+"\"].toSVG()")
	compareXML(t, goXML, nodeXML)

	goXML = octi.ToSVG(nil)
	nodeXML = nodeCmd(t, "require('octicons')[\""+name+"\"].toSVG()")
	compareXML(t, goXML, nodeXML)

	goXML = octi.ToSVGUse(nil)
	nodeXML = nodeCmd(t, "require('octicons')[\""+name+"\"].toSVGUse()")
	compareXML(t, goXML, nodeXML)

	goXML = octi.ToSVGUse(nil)
	nodeXML = nodeCmd(t, "require('octicons')[\""+name+"\"].toSVGUse()")
	compareXML(t, goXML, nodeXML)
}

func nodeCmd(t *testing.T, command string) string {
	out, err := exec.Command("node", "-p", command).Output()
	if err != nil {
		t.Fatal(err)
	}

	return string(out)
}

func compareXML(t *testing.T, xml1 string, xml2 string) {
	assertEqual(t, canonXML(t, xml1), canonXML(t, xml2))
}

func canonXML(t *testing.T, xml string) string {
	cmd := exec.Command("xmllint", "-c14n", "-")
	cmdIn, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}

	cmd.Start()
	cmdIn.Write([]byte(xml))
	cmdIn.Close()
	out, err := ioutil.ReadAll(cmdOut)
	if err != nil {
		t.Fatal(err)
	}
	cmd.Wait()

	return string(out)
}

func assertEqual(t *testing.T, exp interface{}, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", exp, act))
}
