//source: https://doc.qt.io/qt-5/qtwidgets-widgets-lineedits-example.html

package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// Char info Block
type Char_Info interface {
	setup()
	getName()
}

type char_info_block struct {
	echoGroup             *widgets.QGroupBox
	echoLabel             *widgets.QLabel
	echoLineEditCharName  *widgets.QLineEdit
	echoLineEditSpellName *widgets.QLineEdit
	echoLayout            *widgets.QGridLayout
}

func create_char_info() *char_info_block {
	v := char_info_block{widgets.NewQGroupBox2("", nil), widgets.NewQLabel2("Basics:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x char_info_block) setup() {
	x.echoLineEditCharName.SetPlaceholderText("Character Name")
	x.echoLineEditSpellName.SetPlaceholderText("Spell Name")
	x.echoLayout.AddWidget2(x.echoLabel, 0, 0, 0)
	x.echoLayout.AddWidget3(x.echoLineEditCharName, 1, 0, 1, 1, 0)
	x.echoLayout.AddWidget3(x.echoLineEditSpellName, 1, 1, 1, 2, 4)
	x.echoGroup.SetLayout(x.echoLayout)
}

func (x char_info_block) getName() string {
	return x.echoLineEditCharName.Text()
}

//Attack block info
type Attack_blocker interface {
	setup()
}

type attack_block struct {
	AttkGroup              *widgets.QGroupBox
	AttkComboBoxTF         *widgets.QComboBox
	AttkLabel              *widgets.QLabel
	AttkPlus               *widgets.QLabel
	AttkComboBoxType       *widgets.QComboBox
	AttkLineEditModifier   *widgets.QLineEdit
	AttkComboBoxProficient *widgets.QComboBox
	AttkProfLabel          *widgets.QLabel
	AttkRangeLabel         *widgets.QLabel
	AttkLineEditRange      *widgets.QLineEdit
	AttkMagicBonusLabel    *widgets.QLabel
	AttkLineEditMagicBonus *widgets.QLineEdit
	AttkCritRangeLabel     *widgets.QLabel
	AttkLineEditCritRange  *widgets.QLineEdit
	attkLayout             *widgets.QGridLayout
}

func create_attack_block() *attack_block {
	v := attack_block{widgets.NewQGroupBox2("", nil), widgets.NewQComboBox(nil), widgets.NewQLabel2("Attack:", nil, 0), widgets.NewQLabel2("+", nil, 0), widgets.NewQComboBox(nil), widgets.NewQLineEdit(nil), widgets.NewQComboBox(nil), widgets.NewQLabel2(" + Proficient", nil, 0), widgets.NewQLabel2("Range:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Magic Bonus:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Crit Range:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x attack_block) setup() {
	x.AttkComboBoxTF.AddItems([]string{"True", "False"})
	x.AttkComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})

	x.AttkLineEditModifier.SetPlaceholderText("0")
	x.AttkLineEditModifier.SetValidator(gui.NewQIntValidator(x.AttkLineEditModifier))
	x.AttkLineEditMagicBonus.SetPlaceholderText("0")
	x.AttkComboBoxProficient.AddItems([]string{"False", "True"})
	x.AttkLineEditRange.SetPlaceholderText("Self (60ft Cone)")
	x.AttkLineEditCritRange.SetTextDefault("20")

	x.attkLayout.AddWidget2(x.AttkLabel, 0, 1, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxTF, 0, 0, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxType, 0, 2, 0)
	x.attkLayout.AddWidget2(x.AttkPlus, 0, 3, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditModifier, 0, 4, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxProficient, 0, 5, 0)
	x.attkLayout.AddWidget2(x.AttkProfLabel, 0, 6, 0)

	x.attkLayout.AddWidget2(x.AttkRangeLabel, 1, 0, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditRange, 1, 1, 0)

	x.attkLayout.AddWidget2(x.AttkMagicBonusLabel, 2, 0, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditMagicBonus, 2, 1, 0)
	x.attkLayout.AddWidget2(x.AttkCritRangeLabel, 2, 2, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditCritRange, 2, 3, 0)

	x.AttkGroup.SetLayout(x.attkLayout)
}

//Damage block

type Damage_blocker interface {
	setup()
}

type damage_block struct {
	Damage1Group          *widgets.QGroupBox
	Damage1ComboBoxTF     *widgets.QComboBox
	Damage1LabelDamage    *widgets.QLabel
	Damage1LineEditDamage *widgets.QLineEdit
	Damage1Plus1          *widgets.QLabel
	Damage1ComboBoxType   *widgets.QComboBox
	Damage1LabelPlus2     *widgets.QLabel
	Damage1LineEditBonus  *widgets.QLineEdit

	Damage1LabelType    *widgets.QLabel
	Damage1LineEditType *widgets.QLineEdit
	Damage1LabelCrit    *widgets.QLabel
	Damage1LineEditCrit *widgets.QLineEdit
	damage1Layout       *widgets.QGridLayout
}

func create_damage_block() *damage_block {
	v := damage_block{widgets.NewQGroupBox2("", nil), widgets.NewQComboBox(nil), widgets.NewQLabel2("Damage:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("+", nil, 0), widgets.NewQComboBox(nil), widgets.NewQLabel2("+", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Type:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Crit:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x damage_block) setup() {
	x.Damage1ComboBoxTF.AddItems([]string{"True", "False"})
	x.Damage1ComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	x.Damage1LineEditDamage.SetPlaceholderText("1d6")
	x.Damage1LineEditBonus.SetPlaceholderText("0")
	x.Damage1LineEditType.SetPlaceholderText("Slashing")
	x.Damage1LineEditCrit.SetPlaceholderText("1d6")

	x.damage1Layout.AddWidget2(x.Damage1ComboBoxTF, 0, 0, 0)
	x.damage1Layout.AddWidget2(x.Damage1LabelDamage, 0, 1, 0)
	x.damage1Layout.AddWidget2(x.Damage1LineEditDamage, 0, 2, 0)
	x.damage1Layout.AddWidget2(x.Damage1Plus1, 0, 3, 0)
	x.damage1Layout.AddWidget2(x.Damage1ComboBoxTF, 0, 4, 0)
	x.damage1Layout.AddWidget2(x.Damage1LabelPlus2, 0, 5, 0)
	x.damage1Layout.AddWidget2(x.Damage1LineEditBonus, 0, 6, 0)

	x.damage1Layout.AddWidget2(x.Damage1LabelType, 1, 0, 0)
	x.damage1Layout.AddWidget2(x.Damage1LineEditType, 1, 1, 0)
	x.damage1Layout.AddWidget2(x.Damage1LabelCrit, 1, 2, 0)
	x.damage1Layout.AddWidget2(x.Damage1LineEditCrit, 1, 3, 0)

	x.Damage1Group.SetLayout(x.damage1Layout)
}

// Saving Throw Box

type saving_throw interface {
	setup()
}

type saving_throw_block struct {
	SavingThrowGroup            *widgets.QGroupBox
	SavingThrowComboBoxTF       *widgets.QComboBox
	SavingThrowLabelSavingThrow *widgets.QLabel
	SavingThrowComboBoxStat     *widgets.QComboBox
	SavingThrowLabelVsDC        *widgets.QLabel
	SavingThrowComboBoxType     *widgets.QComboBox

	SavingThrowLabelEffect *widgets.QLabel
	SavingThrowLineEdit    *widgets.QLineEdit
	SavingThrowLayout      *widgets.QGridLayout
}

func create_saving_throw_block() *saving_throw_block {
	v := saving_throw_block{widgets.NewQGroupBox2("", nil), widgets.NewQComboBox(nil), widgets.NewQLabel2("Saving Throw:", nil, 0), widgets.NewQComboBox(nil), widgets.NewQLabel2("vs DC:", nil, 0), widgets.NewQComboBox(nil), widgets.NewQLabel2("Save Effect:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x saving_throw_block) setup() {
	x.SavingThrowComboBoxTF.AddItems([]string{"False", "True"})
	x.SavingThrowComboBoxStat.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA"})
	x.SavingThrowComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	x.SavingThrowLineEdit.SetPlaceholderText("Half Damage")

	x.SavingThrowLayout.AddWidget2(x.SavingThrowComboBoxTF, 0, 0, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowLabelSavingThrow, 0, 1, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowComboBoxStat, 0, 2, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowLabelVsDC, 0, 3, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowComboBoxType, 0, 4, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowLabelEffect, 1, 0, 0)
	x.SavingThrowLayout.AddWidget2(x.SavingThrowLineEdit, 1, 1, 0)

	x.SavingThrowGroup.SetLayout(x.SavingThrowLayout)
}

type d20_macro interface {
	setup()
}

type d20_macro_block struct {
	d20_attack_block  *attack_block
	d20_damage_block1 *damage_block
	d20_damage_block2 *damage_block
	d20_saving_block  *saving_throw_block
	Layout            *widgets.QGridLayout
}

func create_d20_macro() *d20_macro_block {
	v := d20_macro_block{create_attack_block(), create_damage_block(), create_damage_block(), create_saving_throw_block(), widgets.NewQGridLayout2()}
	return &v
}

func (x d20_macro_block) setup() {
	x.d20_attack_block.setup()
	x.d20_damage_block1.setup()
	x.d20_damage_block2.setup()
	x.d20_saving_block.setup()

	x.Layout.AddWidget2(x.d20_attack_block.AttkGroup, 1, 0, 0)
	x.Layout.AddWidget2(x.d20_damage_block1.Damage1Group, 2, 0, 0)
	x.Layout.AddWidget2(x.d20_damage_block2.Damage1Group, 3, 0, 0)
	x.Layout.AddWidget2(x.d20_saving_block.SavingThrowGroup, 4, 0, 0)
}
func (x d20_macro_block) getName() string {
	return "Meow"
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	t := create_d20_macro()
	t.setup()

	v := create_d20_macro()
	v.setup()

	x := create_char_info()
	x.setup()

	var myTabWidget *widgets.QTabWidget
	myTabWidget = widgets.NewQTabWidget(nil)

	var myFirstTabPage *widgets.QWidget
	myFirstTabPage = widgets.NewQWidget(nil, 0)
	var mySecondTabPage *widgets.QWidget
	mySecondTabPage = widgets.NewQWidget(nil, 0)

	//echoComboBox.ConnectCurrentIndexChanged(func(index int) { echoChanged(echoLineEditCharName, index) })

	// Button

	var (
		buttonGroup = widgets.NewQGroupBox2("", nil)
		button      = widgets.NewQPushButton2("Generate", nil)
	)

	var ButtonLayout = widgets.NewQGridLayout2()
	ButtonLayout.AddWidget2(button, 0, 0, 0)
	buttonGroup.SetLayout(ButtonLayout)

	button.ConnectClicked(func(thing bool) { generate_macro(thing, x, t, v) })

	myFirstTabPage.SetLayout(v.Layout)
	mySecondTabPage.SetLayout(t.Layout)

	var mainWindowLayout = widgets.NewQGridLayout2()
	mainWindowLayout.AddWidget2(x.echoGroup, 0, 0, 0)
	mainWindowLayout.AddWidget2(myTabWidget, 1, 0, 0)
	mainWindowLayout.AddWidget2(buttonGroup, 1, 1, 0)

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

func generate_macro(thing bool, character *char_info_block, block1 *d20_macro_block, block2 *d20_macro_block) {
	if thing == false {

		var macro string = ""

		macro += "!&{template:atkdmg} "
		macro += "{{mod=@{" + character.getName() + "|spell_attack_bonus}}}"

		macro += "{{rname=test}} "
		macro += "{{r1=[[1d20+@{" + character.getName() + "|spell_attack_bonus}]]}} "
		macro += "{{normal=1}} {{attack=1}} "
		macro += "{{range=20ft}} "
		macro += "{{damage=1}} {{dmg1flag=1}} "
		macro += "{{dmg1=[[1d6]]}}"
		macro += "{{dmg1type=1d6}} {{spelllevel=0}} "
		macro += "}"

		fmt.Println(macro)

	}
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
