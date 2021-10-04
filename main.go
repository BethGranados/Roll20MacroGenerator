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
	echoGroup            *widgets.QGroupBox
	echoLabel            *widgets.QLabel
	echoLineEditCharName *widgets.QLineEdit
	echoLayout           *widgets.QGridLayout
}

func create_char_info() *char_info_block {
	v := char_info_block{widgets.NewQGroupBox2("", nil), widgets.NewQLabel2("Basics:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x char_info_block) setup() {
	x.echoLineEditCharName.SetPlaceholderText("Character Name")

	x.echoLayout.AddWidget2(x.echoLabel, 0, 0, 0)
	x.echoLayout.AddWidget3(x.echoLineEditCharName, 1, 0, 1, 1, 0)
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
	echoLineEditSpellName  *widgets.QLineEdit
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
	v := attack_block{widgets.NewQGroupBox2("", nil), widgets.NewQLineEdit(nil), widgets.NewQComboBox(nil), widgets.NewQLabel2("Attack:", nil, 0), widgets.NewQLabel2("+", nil, 0), widgets.NewQComboBox(nil), widgets.NewQLineEdit(nil), widgets.NewQComboBox(nil), widgets.NewQLabel2(" + Proficient", nil, 0), widgets.NewQLabel2("Range:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Magic Bonus:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQLabel2("Crit Range:", nil, 0), widgets.NewQLineEdit(nil), widgets.NewQGridLayout2()}
	return &v
}

func (x attack_block) setup() {

	x.AttkComboBoxTF.AddItems([]string{"True", "False"})
	x.AttkComboBoxType.AddItems([]string{"STR", "DEX", "CON", "INT", "WIS", "CHA", "SPELL", "-"})
	x.echoLineEditSpellName.SetPlaceholderText("Action Name")

	x.AttkLineEditModifier.SetPlaceholderText("0")
	x.AttkLineEditModifier.SetValidator(gui.NewQIntValidator(x.AttkLineEditModifier))
	x.AttkLineEditMagicBonus.SetPlaceholderText("0")
	x.AttkComboBoxProficient.AddItems([]string{"False", "True"})
	x.AttkLineEditRange.SetPlaceholderText("Self (60ft Cone)")
	x.AttkLineEditCritRange.SetTextDefault("20")

	x.attkLayout.AddWidget2(x.AttkLabel, 0, 0, 0)
	x.attkLayout.AddWidget2(x.echoLineEditSpellName, 0, 1, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxTF, 1, 0, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxType, 1, 2, 0)
	x.attkLayout.AddWidget2(x.AttkPlus, 1, 3, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditModifier, 1, 4, 0)
	x.attkLayout.AddWidget2(x.AttkComboBoxProficient, 1, 5, 0)
	x.attkLayout.AddWidget2(x.AttkProfLabel, 1, 6, 0)

	x.attkLayout.AddWidget2(x.AttkRangeLabel, 2, 0, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditRange, 2, 1, 0)

	x.attkLayout.AddWidget2(x.AttkMagicBonusLabel, 3, 0, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditMagicBonus, 3, 1, 0)
	x.attkLayout.AddWidget2(x.AttkCritRangeLabel, 3, 2, 0)
	x.attkLayout.AddWidget2(x.AttkLineEditCritRange, 3, 3, 0)

	x.AttkGroup.SetLayout(x.attkLayout)
}

func (x attack_block) getMoveName() string {
	return x.echoLineEditSpellName.Text()
}

func (x attack_block) getModType() (bool, string) {
	switch x.AttkComboBoxType.CurrentText() {
	case "STR":
		return true, "strength_mod"
	case "DEX":
		return true, "dexterity_mod"
	case "CON":
		return true, "constitution_mod"
	case "INT":
		return true, "intelligence_mod"
	case "WIS":
		return true, "wisdom_mod"
	case "CHA":
		return true, "charisma_mod"
	case "SPELL":
		return true, "spell_attack_bonus"
	case "-":
		return false, ""
	default:
		return false, ""
	}
}

func (x attack_block) isProficient() bool {
	switch x.AttkComboBoxProficient.CurrentText() {
	case "True":
		return true
	case "False":
		return false
	}
	return false
}

func (x attack_block) getBonusMod() (bool, string) {
	mod := x.AttkLineEditModifier.Text()
	if mod != "" {
		return true, mod
	}
	return false, ""
}

func (x attack_block) getRange() string {
	return x.AttkLineEditRange.Text()
}

func (x attack_block) getMagicBonus() string {
	return x.AttkMagicBonusLabel.Text()
}

func (x attack_block) getCritRange() string {
	return x.AttkCritRangeLabel.Text()
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

func (x damage_block) isEnabled() bool {
	switch x.Damage1ComboBoxTF.CurrentText() {
	case "True":
		return true
	case "False":
		return false
	}
	return false
}

func (x damage_block) getBonusMod() string {
	return x.Damage1LineEditBonus.Text()
}

func (x damage_block) getDamageType() string {
	return x.Damage1LineEditType.Text()
}

func (x damage_block) getCrit() string {
	return x.Damage1LineEditCrit.Text()
}

func (x damage_block) getBonusType() string {
	return x.Damage1LabelType.Text()
}

func (x damage_block) getDamage() string {
	return x.Damage1LineEditDamage.Text()
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

func (x saving_throw_block) isActive() bool {
	switch x.SavingThrowComboBoxTF.CurrentText() {
	case "True":
		return true
	case "False":
		return false
	}
	return false
}

func (x saving_throw_block) getThrowStat() (bool, string) {
	switch x.SavingThrowComboBoxStat.CurrentText() {
	case "STR":
		return true, "strength_mod"
	case "DEX":
		return true, "dexterity_mod"
	case "CON":
		return true, "constitution_mod"
	case "INT":
		return true, "intelligence_mod"
	case "WIS":
		return true, "wisdom_mod"
	case "CHA":
		return true, "charisma_mod"
	default:
		return false, ""
	}
}

func (x saving_throw_block) getSavingThrowEffect() string {
	return x.SavingThrowLineEdit.Text()
}

func (x saving_throw_block) getThrowType() (bool, string) {
	switch x.SavingThrowComboBoxType.CurrentText() {
	case "STR":
		return true, "strength_mod"
	case "DEX":
		return true, "dexterity_mod"
	case "CON":
		return true, "constitution_mod"
	case "INT":
		return true, "intelligence_mod"
	case "WIS":
		return true, "wisdom_mod"
	case "CHA":
		return true, "charisma_mod"
	case "SPELL":
		return true, "spell_attack_bonus"
	case "-":
		return false, ""
	default:
		return false, ""
	}
}

// Put it all together...

type d20_macro interface {
	setup()
	generate_macro()
}

type d20_macro_block struct {
	character         *char_info_block
	d20_attack_block  *attack_block
	d20_damage_block1 *damage_block
	d20_damage_block2 *damage_block
	d20_saving_block  *saving_throw_block
	Layout            *widgets.QGridLayout
}

func create_d20_macro(info *char_info_block) *d20_macro_block {
	v := d20_macro_block{info, create_attack_block(), create_damage_block(), create_damage_block(), create_saving_throw_block(), widgets.NewQGridLayout2()}
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

func (x d20_macro_block) generate_macro() string {
	moveName := x.d20_attack_block.getMoveName()
	definedType, moveType := x.d20_attack_block.getModType()
	definedMod, modifier := x.d20_attack_block.getBonusMod()
	//proficient := x.d20_attack_block.isProficient()
	atkRange := x.d20_attack_block.getRange()
	//magicBonus := x.d20_attack_block.getMagicBonus()
	//Crit := x.d20_attack_block.getCritRange()

	//cb := "}"
	cb := "&#125;"

	macroText := ""

	macroText += "&{template:atkdmg" + cb + "{{mod="

	if definedType {
		macroText += "@{" + x.character.getName() + "|" + moveType + "}"
	}

	if definedMod {
		macroText += " + " + modifier
	}

	macroText += cb + cb + "{{rname=" + moveName + cb + cb + "{{r1=[[1d20"

	if definedType {
		macroText += "+@{" + x.character.getName() + "|" + moveType + "}"
	}

	if definedMod {
		macroText += " + " + modifier
	}

	macroText += "]]" + cb + cb + "{{normal=1" + cb + cb + "{{attack=1" + cb + cb + "{{range=" + atkRange + cb + cb

	if x.d20_damage_block1.isEnabled() {
		macroText += "{{damage=1" + cb + cb + "{{dmg1flag=1" + cb + cb + "{{dmg1=[[" + x.d20_damage_block1.getDamage()
		macroText += "]]" + cb + cb + "{{dmg1type=" + x.d20_damage_block1.getBonusType() + cb + cb
	} else {
		macroText += "{{damage=0" + cb + cb + "{{dmg1flag=" + cb + cb + "{{dmg1=0" + cb + cb + "{{dmg1type=" + cb + cb
	}

	if x.d20_damage_block2.isEnabled() {
		macroText += "{{damage=2" + cb + cb + "{{dmg2flag=1" + cb + cb + "{{dmg2=[[" + x.d20_damage_block2.getDamage() + "]+" + cb + cb + "{{dmg2type=" + cb + cb
	} else {
		macroText += "{{damage=0" + cb + cb + "{{dmg2flag=" + cb + cb + "{{dmg2=0" + cb + cb + "{{dmg2type=" + cb + cb
	}

	macroText += "{{spelllevel=0" + cb + cb

	return macroText

	//return "!&{template:atkdmg}{{mod=@{" + x.character.getName() + "|" + moveType + "}" + modifier + "}}{{rname=" + moveName + "}} {{r1=[[1d20+@{" +
	//x.character.getName() + "|" + moveType + "}  ]]}} {{normal=1}} {{attack=1}} {{range=" + atkRange + "}}{{damage=1}} {{dmg1flag=1}}{{dmg1=[[1d6]]}}{{dmg1type=1d6}} {{spelllevel=0}}}"
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	blockList := []*d20_macro_block{}

	x := create_char_info()
	x.setup()

	t := create_d20_macro(x)
	t.setup()

	blockList = append(blockList, t)

	v := create_d20_macro(x)
	v.setup()

	blockList = append(blockList, v)

	var myTabWidget *widgets.QTabWidget
	myTabWidget = widgets.NewQTabWidget(nil)

	var myFirstTabPage *widgets.QWidget
	myFirstTabPage = widgets.NewQWidget(nil, 0)
	var mySecondTabPage *widgets.QWidget
	mySecondTabPage = widgets.NewQWidget(nil, 0)

	//echoComboBox.ConnectCurrentIndexChanged(func(index int) { echoChanged(echoLineEditCharName, index) })

	// Button

	var (
		buttonGroup    = widgets.NewQGroupBox2("", nil)
		generateButton = widgets.NewQPushButton2("Generate", nil)
	)

	var ButtonLayout = widgets.NewQGridLayout2()
	ButtonLayout.AddWidget2(generateButton, 0, 0, 0)
	buttonGroup.SetLayout(ButtonLayout)

	//generateButton.ConnectClicked(func(thing bool) { generate_macro_full(thing, x, blockList[0], blockList[1]) })

	generateButton.ConnectClicked(func(thing bool) { generate_macro_full(thing, x, blockList) })

	myFirstTabPage.SetLayout(blockList[0].Layout)
	mySecondTabPage.SetLayout(blockList[1].Layout)

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

func generate_macro_full(thing bool, character *char_info_block, blockList []*d20_macro_block) {
	if thing == false {
		/*fmt.Printf("?{Which Option?")
		fmt.Printf(block1.generate_macro())
		fmt.Printf("|Option2,\n")
		fmt.Printf(block2.generate_macro())
		fmt.Printf("}")*/

		for _, j := range blockList {
			fmt.Printf("|Option1,\n")
			fmt.Printf(j.generate_macro())
		}
		fmt.Printf("}")
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
