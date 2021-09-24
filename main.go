//source: https://doc.qt.io/qt-5/qtwidgets-widgets-lineedits-example.html

package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type attack_block struct {
	echoGroup             *widgets.QGroupBox
	echoLabel             *widgets.QLabel
	echoLineEditCharName  *widgets.QLineEdit
	echoLineEditSpellName *widgets.QLineEdit
	echoLayout            *widgets.QGridLayout
}

func (x attack_block) Error() {
	fmt.Println("Invalid query entered")
}

func create_attack_block() *attack_block {
	v := attack_block{widgets.NewQGroupBox2("", nil), widgets.NewQLabel2("Basics:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x attack_block) setup() {
	x.echoLineEditCharName.SetPlaceholderText("Character Name")
	x.echoLineEditSpellName.SetPlaceholderText("Spell Name")
	x.echoLayout.AddWidget2(x.echoLabel, 0, 0, 0)
	x.echoLayout.AddWidget3(x.echoLineEditCharName, 1, 0, 1, 1, 0)
	x.echoLayout.AddWidget3(x.echoLineEditSpellName, 1, 1, 1, 2, 4)
	x.echoGroup.SetLayout(x.echoLayout)
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	w := create_attack_block()
	w.setup()

	var myTabWidget *widgets.QTabWidget
	myTabWidget = widgets.NewQTabWidget(nil)

	var myFirstTabPage *widgets.QWidget
	myFirstTabPage = widgets.NewQWidget(nil, 0)
	var mySecondTabPage *widgets.QWidget
	mySecondTabPage = widgets.NewQWidget(nil, 0)

	// Echo Group Stuff
	var (
		echoGroup             = widgets.NewQGroupBox2("", nil)
		echoLabel             = widgets.NewQLabel2("Basics:", nil, 0)
		echoLineEditCharName  = widgets.NewQLineEdit(nil)
		echoLineEditSpellName = widgets.NewQLineEdit(nil)
	)

	//echoComboBox.AddItems([]string{"STR", "DEX", "CHA", "CON", "INT", "WIS", "SPELL"})
	echoLineEditCharName.SetPlaceholderText("Character Name")
	echoLineEditSpellName.SetPlaceholderText("Spell Name")

	var echoLayout = widgets.NewQGridLayout2()
	echoLayout.AddWidget2(echoLabel, 0, 0, 0)
	//echoLayout.AddWidget2(echoComboBox, 0, 1, 0)
	echoLayout.AddWidget3(echoLineEditCharName, 1, 0, 1, 1, 0)
	echoLayout.AddWidget3(echoLineEditSpellName, 1, 1, 1, 2, 4)
	echoGroup.SetLayout(echoLayout)
	//myFirstTabPage.SetLayout(echoLayout)

	//echoComboBox.ConnectCurrentIndexChanged(func(index int) { echoChanged(echoLineEditCharName, index) })

	// Attack Group Stuff
	var (
		AttkGroup              = widgets.NewQGroupBox2("", nil)
		AttkComboBoxTF         = widgets.NewQComboBox(nil)
		AttkLabel              = widgets.NewQLabel2("Attack:", nil, 0)
		AttkPlus               = widgets.NewQLabel2("+", nil, 0)
		AttkComboBoxType       = widgets.NewQComboBox(nil)
		AttkLineEditModifier   = widgets.NewQLineEdit(nil)
		AttkComboBoxProficient = widgets.NewQComboBox(nil)
		AttkProfLabel          = widgets.NewQLabel2(" + Proficient", nil, 0)

		AttkRangeLabel    = widgets.NewQLabel2("Range:", nil, 0)
		AttkLineEditRange = widgets.NewQLineEdit(nil)

		AttkMagicBonusLabel    = widgets.NewQLabel2("Magic Bonus:", nil, 0)
		AttkLineEditMagicBonus = widgets.NewQLineEdit(nil)
		AttkCritRangeLabel     = widgets.NewQLabel2("Crit Range:", nil, 0)
		AttkLineEditCritRange  = widgets.NewQLineEdit(nil)
	)

	AttkComboBoxTF.AddItems([]string{"True", "False"})
	AttkComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})

	AttkLineEditModifier.SetPlaceholderText("0")
	AttkLineEditModifier.SetValidator(gui.NewQIntValidator(AttkLineEditModifier))
	AttkLineEditMagicBonus.SetPlaceholderText("0")
	AttkComboBoxProficient.AddItems([]string{"False", "True"})
	AttkLineEditRange.SetPlaceholderText("Self (60ft Cone)")
	//AttkLineEditCritRange.SetPlaceholderText("20")
	AttkLineEditCritRange.SetTextDefault("20")

	var attkLayout = widgets.NewQGridLayout2()
	attkLayout.AddWidget2(AttkLabel, 0, 1, 0)
	attkLayout.AddWidget2(AttkComboBoxTF, 0, 0, 0)
	attkLayout.AddWidget2(AttkComboBoxType, 0, 2, 0)
	attkLayout.AddWidget2(AttkPlus, 0, 3, 0)
	attkLayout.AddWidget2(AttkLineEditModifier, 0, 4, 0)
	attkLayout.AddWidget2(AttkComboBoxProficient, 0, 5, 0)
	attkLayout.AddWidget2(AttkProfLabel, 0, 6, 0)

	attkLayout.AddWidget2(AttkRangeLabel, 1, 0, 0)
	attkLayout.AddWidget2(AttkLineEditRange, 1, 1, 0)

	attkLayout.AddWidget2(AttkMagicBonusLabel, 2, 0, 0)
	attkLayout.AddWidget2(AttkLineEditMagicBonus, 2, 1, 0)
	attkLayout.AddWidget2(AttkCritRangeLabel, 2, 2, 0)
	attkLayout.AddWidget2(AttkLineEditCritRange, 2, 3, 0)

	AttkGroup.SetLayout(attkLayout)

	//Damage 1 Stuff
	var (
		Damage1Group          = widgets.NewQGroupBox2("", nil)
		Damage1ComboBoxTF     = widgets.NewQComboBox(nil)
		Damage1LabelDamage    = widgets.NewQLabel2("Damage:", nil, 0)
		Damage1LineEditDamage = widgets.NewQLineEdit(nil)
		Damage1Plus1          = widgets.NewQLabel2("+", nil, 0)
		Damage1ComboBoxType   = widgets.NewQComboBox(nil)
		Damage1LabelPlus2     = widgets.NewQLabel2("+", nil, 0)
		Damage1LineEditBonus  = widgets.NewQLineEdit(nil)

		Damage1LabelType    = widgets.NewQLabel2("Type:", nil, 0)
		Damage1LineEditType = widgets.NewQLineEdit(nil)
		Damage1LabelCrit    = widgets.NewQLabel2("Crit:", nil, 0)
		Damage1LineEditCrit = widgets.NewQLineEdit(nil)
	)

	Damage1ComboBoxTF.AddItems([]string{"True", "False"})
	Damage1ComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	Damage1LineEditDamage.SetPlaceholderText("1d6")
	Damage1LineEditBonus.SetPlaceholderText("0")
	Damage1LineEditType.SetPlaceholderText("Slashing")
	Damage1LineEditCrit.SetPlaceholderText("1d6")

	var damage1Layout = widgets.NewQGridLayout2()
	damage1Layout.AddWidget2(Damage1ComboBoxTF, 0, 0, 0)
	damage1Layout.AddWidget2(Damage1LabelDamage, 0, 1, 0)
	damage1Layout.AddWidget2(Damage1LineEditDamage, 0, 2, 0)
	damage1Layout.AddWidget2(Damage1Plus1, 0, 3, 0)
	damage1Layout.AddWidget2(Damage1ComboBoxTF, 0, 4, 0)
	damage1Layout.AddWidget2(Damage1LabelPlus2, 0, 5, 0)
	damage1Layout.AddWidget2(Damage1LineEditBonus, 0, 6, 0)

	damage1Layout.AddWidget2(Damage1LabelType, 1, 0, 0)
	damage1Layout.AddWidget2(Damage1LineEditType, 1, 1, 0)
	damage1Layout.AddWidget2(Damage1LabelCrit, 1, 2, 0)
	damage1Layout.AddWidget2(Damage1LineEditCrit, 1, 3, 0)

	Damage1Group.SetLayout(damage1Layout)

	//Damage 2 Stuff
	var (
		Damage2Group          = widgets.NewQGroupBox2("", nil)
		Damage2ComboBoxTF     = widgets.NewQComboBox(nil)
		Damage2LabelDamage    = widgets.NewQLabel2("Damage:", nil, 0)
		Damage2LineEditDamage = widgets.NewQLineEdit(nil)
		Damage2Plus1          = widgets.NewQLabel2("+", nil, 0)
		Damage2ComboBoxType   = widgets.NewQComboBox(nil)
		Damage2LabelPlus2     = widgets.NewQLabel2("+", nil, 0)
		Damage2LineEditBonus  = widgets.NewQLineEdit(nil)

		Damage2LabelType    = widgets.NewQLabel2("Type:", nil, 0)
		Damage2LineEditType = widgets.NewQLineEdit(nil)
		Damage2LabelCrit    = widgets.NewQLabel2("Crit:", nil, 0)
		Damage2LineEditCrit = widgets.NewQLineEdit(nil)
	)

	Damage2ComboBoxTF.AddItems([]string{"False", "True"})
	Damage2ComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	Damage2LineEditDamage.SetPlaceholderText("1d6")
	Damage2LineEditBonus.SetPlaceholderText("0")
	Damage2LineEditType.SetPlaceholderText("Slashing")
	Damage2LineEditCrit.SetPlaceholderText("1d6")

	var Damage2Layout = widgets.NewQGridLayout2()
	Damage2Layout.AddWidget2(Damage2ComboBoxTF, 0, 0, 0)
	Damage2Layout.AddWidget2(Damage2LabelDamage, 0, 1, 0)
	Damage2Layout.AddWidget2(Damage2LineEditDamage, 0, 2, 0)
	Damage2Layout.AddWidget2(Damage2Plus1, 0, 3, 0)
	Damage2Layout.AddWidget2(Damage2ComboBoxTF, 0, 4, 0)
	Damage2Layout.AddWidget2(Damage2LabelPlus2, 0, 5, 0)
	Damage2Layout.AddWidget2(Damage2LineEditBonus, 0, 6, 0)

	Damage2Layout.AddWidget2(Damage2LabelType, 1, 0, 0)
	Damage2Layout.AddWidget2(Damage2LineEditType, 1, 1, 0)
	Damage2Layout.AddWidget2(Damage2LabelCrit, 1, 2, 0)
	Damage2Layout.AddWidget2(Damage2LineEditCrit, 1, 3, 0)

	Damage2Group.SetLayout(Damage2Layout)

	//Saving Throw
	var (
		SavingThrowGroup            = widgets.NewQGroupBox2("", nil)
		SavingThrowComboBoxTF       = widgets.NewQComboBox(nil)
		SavingThrowLabelSavingThrow = widgets.NewQLabel2("Saving Throw:", nil, 0)
		SavingThrowComboBoxStat     = widgets.NewQComboBox(nil)
		SavingThrowLabelVsDC        = widgets.NewQLabel2("vs DC:", nil, 0)
		SavingThrowComboBoxType     = widgets.NewQComboBox(nil)

		SavingThrowLabelEffect = widgets.NewQLabel2("Save Effect:", nil, 0)
		SavingThrowLineEdit    = widgets.NewQLineEdit(nil)
	)

	SavingThrowComboBoxTF.AddItems([]string{"False", "True"})
	SavingThrowComboBoxStat.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA"})
	SavingThrowComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	SavingThrowLineEdit.SetPlaceholderText("Half Damage")

	var SavingThrowLayout = widgets.NewQGridLayout2()
	//SavingThrowLayout.AddWidget2(SavingThrowGroup, 0, 0, 0)
	SavingThrowLayout.AddWidget2(SavingThrowComboBoxTF, 0, 0, 0)
	SavingThrowLayout.AddWidget2(SavingThrowLabelSavingThrow, 0, 1, 0)
	SavingThrowLayout.AddWidget2(SavingThrowComboBoxStat, 0, 2, 0)
	SavingThrowLayout.AddWidget2(SavingThrowLabelVsDC, 0, 3, 0)
	SavingThrowLayout.AddWidget2(SavingThrowComboBoxType, 0, 4, 0)
	SavingThrowLayout.AddWidget2(SavingThrowLabelEffect, 1, 0, 0)
	SavingThrowLayout.AddWidget2(SavingThrowLineEdit, 1, 1, 0)

	SavingThrowGroup.SetLayout(SavingThrowLayout)

	// Button

	var (
		buttonGroup = widgets.NewQGroupBox2("", nil)
		button      = widgets.NewQPushButton2("Click me", nil)
	)

	var ButtonLayout = widgets.NewQGridLayout2()
	ButtonLayout.AddWidget2(button, 0, 0, 0)
	buttonGroup.SetLayout(ButtonLayout)

	// Put them together

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget2(echoGroup, 1, 0, 0)
	layout.AddWidget2(AttkGroup, 2, 0, 0)
	layout.AddWidget2(Damage1Group, 3, 0, 0)
	layout.AddWidget2(buttonGroup, 4, 3, 0)
	layout.AddWidget2(Damage2Group, 4, 0, 0)
	layout.AddWidget2(SavingThrowGroup, 5, 0, 0)
	layout.AddWidget2(w.echoGroup, 6, 0, 0)

	myFirstTabPage.SetLayout(layout)

	var mainWindowLayout = widgets.NewQVBoxLayout()
	mainWindowLayout.AddWidget(myTabWidget, 0, core.Qt__AlignCenter)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Line Edits")

	myTabWidget.AddTab(myFirstTabPage, "first")
	myTabWidget.AddTab(mySecondTabPage, "second")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(mainWindowLayout)

	window.SetCentralWidget(centralWidget)

	//tabSwitcherComboBox.ConnectCurrentIndexChanged(func(index int) { echoChanged(window, centralWidget, centralWidget2, index) })

	window.Show()

	widgets.QApplication_Exec()
}

func echoChanged(window *widgets.QMainWindow, centralWidget1 *widgets.QWidget, centralWidget2 *widgets.QWidget, index int) {
	switch index {
	case 0:
		{
			window.SetCentralWidget(centralWidget1)
		}

	case 1:
		{
			window.SetCentralWidget(centralWidget2)
		}

	}
}
