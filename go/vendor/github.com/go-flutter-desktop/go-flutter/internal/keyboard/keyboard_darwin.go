package keyboard

import "github.com/go-gl/glfw/v3.3/glfw"

// DetectTextInputDoneMod returns true if the modifiers pressed
// indicate the typed text can be committed
func DetectTextInputDoneMod(mods glfw.ModifierKey) bool {
	return mods&glfw.ModSuper != 0
}

// platfromNormalize normalizes for macos
func (e *Event) platfromNormalize(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	macosMods := ToMacOSModifiers(mods)
	if val, ok := AsMacOSModifiers(key); ok {
		// On GLFW, the "modifiers" keycode is the state as it is BEFORE this event
		// happened, not AFTER, like every other platform.
		macosMods = val | int(macosMods)
	}
	e.Keymap = "macos"
	e.Modifiers = macosMods
	e.KeyCode = ToMacOSKeyCode(key, scancode)
	e.Characters = e.Character
	e.CharactersIgnoringModifiers = e.Character

}

// Modifier key masks pulled from flutter/packages/flutter/lib/src/services/raw_keyboard_macos.dart
// URL: https://github.com/flutter/flutter/blob/3e63411256cc88afc48044aa5ea06c5c9c6a6846/packages/flutter/lib/src/services/raw_keyboard_macos.dart#L241
const (
	modifierControl    = 0x40000
	modifierShift      = 0x20000
	modifierOption     = 0x80000
	modifierCommand    = 0x100000
	modifierCapsLock   = 0x10000
	modifierNumericPad = 0x200000
)

var modifierKeytoMods = map[glfw.Key]int{
	glfw.KeyLeftControl:  modifierControl,
	glfw.KeyLeftShift:    modifierShift,
	glfw.KeyLeftAlt:      modifierOption,
	glfw.KeyLeftSuper:    modifierCommand,
	glfw.KeyRightControl: modifierControl,
	glfw.KeyRightShift:   modifierShift,
	glfw.KeyRightAlt:     modifierOption,
	glfw.KeyRightSuper:   modifierCommand,
	glfw.KeyCapsLock:     modifierCapsLock,
	glfw.KeyNumLock:      modifierNumericPad,
}

// AsMacOSModifiers translate the keycode to the ModifierKey
func AsMacOSModifiers(keycode glfw.Key) (int, bool) {
	val, ok := modifierKeytoMods[keycode]
	return val, ok
}

// ToMacOSModifiers takes a glfw ModifierKey and return his MacOS equivalent
// as defined in https://github.com/flutter/flutter/blob/3e63411256cc88afc48044aa5ea06c5c9c6a6846/packages/flutter/lib/src/services/raw_keyboard_macos.dart#L241
func ToMacOSModifiers(mods glfw.ModifierKey) (macOSmods int) {
	if mods&glfw.ModControl != 0 {
		macOSmods |= modifierControl
	}
	if mods&glfw.ModShift != 0 {
		macOSmods |= modifierShift
	}
	if mods&glfw.ModAlt != 0 {
		macOSmods |= modifierOption
	}
	if mods&glfw.ModSuper != 0 {
		macOSmods |= modifierCommand
	}
	return macOSmods
}

// ToMacOSKeyCode takes a glfw keyCode and return his MacOS equivalent
/// MacOS doesn't provide a scan code, but a virtual keycode to represent a
//  physical key.
//  We first try to convert the physical key to the virtual keycode and then
//  return the MacOS keycode version of this key.
//  If we fail to get the virtual keycode, map the physical GLFW keycode to the
//  MacOS on.
func ToMacOSKeyCode(keycode glfw.Key, scancode int) int {

	// Map virtual key to a intermediate knownLogicalKeys, than maps the
	// knownLogicalKeys to the macOsToPhysicalKey.
	// This takes into account keyboard mapping. (Azerty keyboard on a Qwerty layout)
	// Example "A" on a Qwerty layout but with a `setxkbmap fr` should return "Q"
	utf8 := glfw.GetKeyName(keycode, scancode)
	if len(utf8) > 0 {
		keyLabel := int([]rune(utf8)[0])
		if val, ok := knownLogicalKeys[keyLabel]; ok {
			if macOSKey, ok := macOsToPhysicalKey[val]; ok {
				return macOSKey
			}
		}
	}

	// If the key is a non-printable on.
	// we cannot use knownLogicalKeys.
	// Example "Backspace".
	if val, ok := glfwToLogicalKey[int(keycode)]; ok {
		if macOSKey, ok := macOsToPhysicalKey[val]; ok {
			return macOSKey
		}
	}

	return 0
}

// knownLogicalKeys pulled from flutter/packages/flutter/lib/src/services/keyboard_key.dart
// URL: https://github.com/flutter/flutter/blob/210f4d83136a2eae773bda471db90ac27676f662/packages/flutter/lib/src/services/keyboard_key.dart#L1667
var knownLogicalKeys = map[int]string{
	0x0100000000:  "none",
	0x0100000010:  "hyper",
	0x0100000011:  "superKey",
	0x0100000013:  "fnLock",
	0x0100000014:  "suspend",
	0x0100000015:  "resume",
	0x0100000016:  "turbo",
	0x0100000017:  "privacyScreenToggle",
	0x0100010082:  "sleep",
	0x0100010083:  "wakeUp",
	0x01000100b5:  "displayToggleIntExt",
	0x0100070000:  "usbReserved",
	0x0100070001:  "usbErrorRollOver",
	0x0100070002:  "usbPostFail",
	0x0100070003:  "usbErrorUndefined",
	0x0000000061:  "keyA",
	0x0000000062:  "keyB",
	0x0000000063:  "keyC",
	0x0000000064:  "keyD",
	0x0000000065:  "keyE",
	0x0000000066:  "keyF",
	0x0000000067:  "keyG",
	0x0000000068:  "keyH",
	0x0000000069:  "keyI",
	0x000000006a:  "keyJ",
	0x000000006b:  "keyK",
	0x000000006c:  "keyL",
	0x000000006d:  "keyM",
	0x000000006e:  "keyN",
	0x000000006f:  "keyO",
	0x0000000070:  "keyP",
	0x0000000071:  "keyQ",
	0x0000000072:  "keyR",
	0x0000000073:  "keyS",
	0x0000000074:  "keyT",
	0x0000000075:  "keyU",
	0x0000000076:  "keyV",
	0x0000000077:  "keyW",
	0x0000000078:  "keyX",
	0x0000000079:  "keyY",
	0x000000007a:  "keyZ",
	0x0000000031:  "digit1",
	0x0000000032:  "digit2",
	0x0000000033:  "digit3",
	0x0000000034:  "digit4",
	0x0000000035:  "digit5",
	0x0000000036:  "digit6",
	0x0000000037:  "digit7",
	0x0000000038:  "digit8",
	0x0000000039:  "digit9",
	0x0000000030:  "digit0",
	0x0100070028:  "enter",
	0x0100070029:  "escape",
	0x010007002a:  "backspace",
	0x010007002b:  "tab",
	0x0000000020:  "space",
	0x000000002d:  "minus",
	0x000000003d:  "equal",
	0x000000005b:  "bracketLeft",
	0x000000005d:  "bracketRight",
	0x000000005c:  "backslash",
	0x000000003b:  "semicolon",
	0x0000000027:  "quote",
	0x0000000060:  "backquote",
	0x000000002c:  "comma",
	0x000000002e:  "period",
	0x000000002f:  "slash",
	0x0100070039:  "capsLock",
	0x010007003a:  "f1",
	0x010007003b:  "f2",
	0x010007003c:  "f3",
	0x010007003d:  "f4",
	0x010007003e:  "f5",
	0x010007003f:  "f6",
	0x0100070040:  "f7",
	0x0100070041:  "f8",
	0x0100070042:  "f9",
	0x0100070043:  "f10",
	0x0100070044:  "f11",
	0x0100070045:  "f12",
	0x0100070046:  "printScreen",
	0x0100070047:  "scrollLock",
	0x0100070048:  "pause",
	0x0100070049:  "insert",
	0x010007004a:  "home",
	0x010007004b:  "pageUp",
	0x010007004c:  "delete",
	0x010007004d:  "end",
	0x010007004e:  "pageDown",
	0x010007004f:  "arrowRight",
	0x0100070050:  "arrowLeft",
	0x0100070051:  "arrowDown",
	0x0100070052:  "arrowUp",
	0x0100070053:  "numLock",
	0x0100070054:  "numpadDivide",
	0x0100070055:  "numpadMultiply",
	0x0100070056:  "numpadSubtract",
	0x0100070057:  "numpadAdd",
	0x0100070058:  "numpadEnter",
	0x0100070059:  "numpad1",
	0x010007005a:  "numpad2",
	0x010007005b:  "numpad3",
	0x010007005c:  "numpad4",
	0x010007005d:  "numpad5",
	0x010007005e:  "numpad6",
	0x010007005f:  "numpad7",
	0x0100070060:  "numpad8",
	0x0100070061:  "numpad9",
	0x0100070062:  "numpad0",
	0x0100070063:  "numpadDecimal",
	0x0100070064:  "intlBackslash",
	0x0100070065:  "contextMenu",
	0x0100070066:  "power",
	0x0100070067:  "numpadEqual",
	0x0100070068:  "f13",
	0x0100070069:  "f14",
	0x010007006a:  "f15",
	0x010007006b:  "f16",
	0x010007006c:  "f17",
	0x010007006d:  "f18",
	0x010007006e:  "f19",
	0x010007006f:  "f20",
	0x0100070070:  "f21",
	0x0100070071:  "f22",
	0x0100070072:  "f23",
	0x0100070073:  "f24",
	0x0100070074:  "open",
	0x0100070075:  "help",
	0x0100070077:  "select",
	0x0100070079:  "again",
	0x010007007a:  "undo",
	0x010007007b:  "cut",
	0x010007007c:  "copy",
	0x010007007d:  "paste",
	0x010007007e:  "find",
	0x010007007f:  "audioVolumeMute",
	0x0100070080:  "audioVolumeUp",
	0x0100070081:  "audioVolumeDown",
	0x0100070085:  "numpadComma",
	0x0100070087:  "intlRo",
	0x0100070088:  "kanaMode",
	0x0100070089:  "intlYen",
	0x010007008a:  "convert",
	0x010007008b:  "nonConvert",
	0x0100070090:  "lang1",
	0x0100070091:  "lang2",
	0x0100070092:  "lang3",
	0x0100070093:  "lang4",
	0x0100070094:  "lang5",
	0x010007009b:  "abort",
	0x01000700a3:  "props",
	0x01000700b6:  "numpadParenLeft",
	0x01000700b7:  "numpadParenRight",
	0x01000700bb:  "numpadBackspace",
	0x01000700d0:  "numpadMemoryStore",
	0x01000700d1:  "numpadMemoryRecall",
	0x01000700d2:  "numpadMemoryClear",
	0x01000700d3:  "numpadMemoryAdd",
	0x01000700d4:  "numpadMemorySubtract",
	0x01000700d7:  "numpadSignChange",
	0x01000700d8:  "numpadClear",
	0x01000700d9:  "numpadClearEntry",
	0x01000700e0:  "controlLeft",
	0x01000700e1:  "shiftLeft",
	0x01000700e2:  "altLeft",
	0x01000700e3:  "metaLeft",
	0x01000700e4:  "controlRight",
	0x01000700e5:  "shiftRight",
	0x01000700e6:  "altRight",
	0x01000700e7:  "metaRight",
	0x01000c0060:  "info",
	0x01000c0061:  "closedCaptionToggle",
	0x01000c006f:  "brightnessUp",
	0x01000c0070:  "brightnessDown",
	0x01000c0072:  "brightnessToggle",
	0x01000c0073:  "brightnessMinimum",
	0x01000c0074:  "brightnessMaximum",
	0x01000c0075:  "brightnessAuto",
	0x01000c0083:  "mediaLast",
	0x01000c008c:  "launchPhone",
	0x01000c008d:  "programGuide",
	0x01000c0094:  "exit",
	0x01000c009c:  "channelUp",
	0x01000c009d:  "channelDown",
	0x01000c00b0:  "mediaPlay",
	0x01000c00b1:  "mediaPause",
	0x01000c00b2:  "mediaRecord",
	0x01000c00b3:  "mediaFastForward",
	0x01000c00b4:  "mediaRewind",
	0x01000c00b5:  "mediaTrackNext",
	0x01000c00b6:  "mediaTrackPrevious",
	0x01000c00b7:  "mediaStop",
	0x01000c00b8:  "eject",
	0x01000c00cd:  "mediaPlayPause",
	0x01000c00cf:  "speechInputToggle",
	0x01000c00e5:  "bassBoost",
	0x01000c0183:  "mediaSelect",
	0x01000c0184:  "launchWordProcessor",
	0x01000c0186:  "launchSpreadsheet",
	0x01000c018a:  "launchMail",
	0x01000c018d:  "launchContacts",
	0x01000c018e:  "launchCalendar",
	0x01000c0192:  "launchApp2",
	0x01000c0194:  "launchApp1",
	0x01000c0196:  "launchInternetBrowser",
	0x01000c019c:  "logOff",
	0x01000c019e:  "lockScreen",
	0x01000c019f:  "launchControlPanel",
	0x01000c01a2:  "selectTask",
	0x01000c01a7:  "launchDocuments",
	0x01000c01ab:  "spellCheck",
	0x01000c01ae:  "launchKeyboardLayout",
	0x01000c01b1:  "launchScreenSaver",
	0x01000c01cb:  "launchAssistant",
	0x01000c01b7:  "launchAudioBrowser",
	0x01000c0201:  "newKey",
	0x01000c0203:  "close",
	0x01000c0207:  "save",
	0x01000c0208:  "print",
	0x01000c0221:  "browserSearch",
	0x01000c0223:  "browserHome",
	0x01000c0224:  "browserBack",
	0x01000c0225:  "browserForward",
	0x01000c0226:  "browserStop",
	0x01000c0227:  "browserRefresh",
	0x01000c022a:  "browserFavorites",
	0x01000c022d:  "zoomIn",
	0x01000c022e:  "zoomOut",
	0x01000c0232:  "zoomToggle",
	0x01000c0279:  "redo",
	0x01000c0289:  "mailReply",
	0x01000c028b:  "mailForward",
	0x01000c028c:  "mailSend",
	0x01000c029d:  "keyboardLayoutSelect",
	0x01000c029f:  "showAllWindows",
	0x010005ff01:  "gameButton1",
	0x010005ff02:  "gameButton2",
	0x010005ff03:  "gameButton3",
	0x010005ff04:  "gameButton4",
	0x010005ff05:  "gameButton5",
	0x010005ff06:  "gameButton6",
	0x010005ff07:  "gameButton7",
	0x010005ff08:  "gameButton8",
	0x010005ff09:  "gameButton9",
	0x010005ff0a:  "gameButton10",
	0x010005ff0b:  "gameButton11",
	0x010005ff0c:  "gameButton12",
	0x010005ff0d:  "gameButton13",
	0x010005ff0e:  "gameButton14",
	0x010005ff0f:  "gameButton15",
	0x010005ff10:  "gameButton16",
	0x010005ff11:  "gameButtonA",
	0x010005ff12:  "gameButtonB",
	0x010005ff13:  "gameButtonC",
	0x010005ff14:  "gameButtonLeft1",
	0x010005ff15:  "gameButtonLeft2",
	0x010005ff16:  "gameButtonMode",
	0x010005ff17:  "gameButtonRight1",
	0x010005ff18:  "gameButtonRight2",
	0x010005ff19:  "gameButtonSelect",
	0x010005ff1a:  "gameButtonStart",
	0x010005ff1b:  "gameButtonThumbLeft",
	0x010005ff1c:  "gameButtonThumbRight",
	0x010005ff1d:  "gameButtonX",
	0x010005ff1e:  "gameButtonY",
	0x010005ff1f:  "gameButtonZ",
	0x0100000012:  "fn",
	0x201000700e1: "shift",
	0x201000700e3: "meta",
	0x201000700e2: "alt",
	0x201000700e0: "control",
}

// glfwToLogicalKey puled fromflutter/lib/src/services/keyboard_maps.dart
// URL: https://github.com/flutter/flutter/blob/ab14307e0c77e4a03ea23ccaa68436d9128a445d/packages/flutter/lib/src/services/keyboard_maps.dart#L1175
var glfwToLogicalKey = map[int]string{
	65:  "keyA",
	66:  "keyB",
	67:  "keyC",
	68:  "keyD",
	69:  "keyE",
	70:  "keyF",
	71:  "keyG",
	72:  "keyH",
	73:  "keyI",
	74:  "keyJ",
	75:  "keyK",
	76:  "keyL",
	77:  "keyM",
	78:  "keyN",
	79:  "keyO",
	80:  "keyP",
	81:  "keyQ",
	82:  "keyR",
	83:  "keyS",
	84:  "keyT",
	85:  "keyU",
	86:  "keyV",
	87:  "keyW",
	88:  "keyX",
	89:  "keyY",
	90:  "keyZ",
	49:  "digit1",
	50:  "digit2",
	51:  "digit3",
	52:  "digit4",
	53:  "digit5",
	54:  "digit6",
	55:  "digit7",
	56:  "digit8",
	57:  "digit9",
	48:  "digit0",
	257: "enter",
	256: "escape",
	259: "backspace",
	258: "tab",
	32:  "space",
	45:  "minus",
	61:  "equal",
	91:  "bracketLeft",
	93:  "bracketRight",
	92:  "backslash",
	59:  "semicolon",
	39:  "quote",
	96:  "backquote",
	44:  "comma",
	46:  "period",
	47:  "slash",
	280: "capsLock",
	290: "f1",
	291: "f2",
	292: "f3",
	293: "f4",
	294: "f5",
	295: "f6",
	296: "f7",
	297: "f8",
	298: "f9",
	299: "f10",
	300: "f11",
	301: "f12",
	283: "printScreen",
	284: "pause",
	260: "insert",
	268: "home",
	266: "pageUp",
	261: "delete",
	269: "end",
	267: "pageDown",
	262: "arrowRight",
	263: "arrowLeft",
	264: "arrowDown",
	265: "arrowUp",
	282: "numLock",
	331: "numpadDivide",
	332: "numpadMultiply",
	334: "numpadAdd",
	335: "numpadEnter",
	321: "numpad1",
	322: "numpad2",
	323: "numpad3",
	324: "numpad4",
	325: "numpad5",
	326: "numpad6",
	327: "numpad7",
	328: "numpad8",
	329: "numpad9",
	320: "numpad0",
	330: "numpadDecimal",
	348: "contextMenu",
	336: "numpadEqual",
	302: "f13",
	303: "f14",
	304: "f15",
	305: "f16",
	306: "f17",
	307: "f18",
	308: "f19",
	309: "f20",
	310: "f21",
	311: "f22",
	312: "f23",
	341: "controlLeft",
	340: "shiftLeft",
	342: "altLeft",
	343: "metaLeft",
	345: "controlRight",
	344: "shiftRight",
	346: "altRight",
	347: "metaRight",
}

// macOsToPhysicalKey puled flutter/lib/src/services/keyboard_maps.dart
// URL: https://github.com/flutter/flutter/blob/ab14307e0c77e4a03ea23ccaa68436d9128a445d/packages/flutter/lib/src/services/keyboard_maps.dart#L1003
// vim cmd reverse map: '<,'>s/\v\s([^:]*):\s*([^,]*),/\2: \1,
var macOsToPhysicalKey = map[string]int{
	"keyA":            0x00000000,
	"keyB":            0x0000000b,
	"keyC":            0x00000008,
	"keyD":            0x00000002,
	"keyE":            0x0000000e,
	"keyF":            0x00000003,
	"keyG":            0x00000005,
	"keyH":            0x00000004,
	"keyI":            0x00000022,
	"keyJ":            0x00000026,
	"keyK":            0x00000028,
	"keyL":            0x00000025,
	"keyM":            0x0000002e,
	"keyN":            0x0000002d,
	"keyO":            0x0000001f,
	"keyP":            0x00000023,
	"keyQ":            0x0000000c,
	"keyR":            0x0000000f,
	"keyS":            0x00000001,
	"keyT":            0x00000011,
	"keyU":            0x00000020,
	"keyV":            0x00000009,
	"keyW":            0x0000000d,
	"keyX":            0x00000007,
	"keyY":            0x00000010,
	"keyZ":            0x00000006,
	"digit1":          0x00000012,
	"digit2":          0x00000013,
	"digit3":          0x00000014,
	"digit4":          0x00000015,
	"digit5":          0x00000017,
	"digit6":          0x00000016,
	"digit7":          0x0000001a,
	"digit8":          0x0000001c,
	"digit9":          0x00000019,
	"digit0":          0x0000001d,
	"enter":           0x00000024,
	"escape":          0x00000035,
	"backspace":       0x00000033,
	"tab":             0x00000030,
	"space":           0x00000031,
	"minus":           0x0000001b,
	"equal":           0x00000018,
	"bracketLeft":     0x00000021,
	"bracketRight":    0x0000001e,
	"backslash":       0x0000002a,
	"semicolon":       0x00000029,
	"quote":           0x00000027,
	"backquote":       0x00000032,
	"comma":           0x0000002b,
	"period":          0x0000002f,
	"slash":           0x0000002c,
	"capsLock":        0x00000039,
	"f1":              0x0000007a,
	"f2":              0x00000078,
	"f3":              0x00000063,
	"f4":              0x00000076,
	"f5":              0x00000060,
	"f6":              0x00000061,
	"f7":              0x00000062,
	"f8":              0x00000064,
	"f9":              0x00000065,
	"f10":             0x0000006d,
	"f11":             0x00000067,
	"f12":             0x0000006f,
	"insert":          0x00000072,
	"home":            0x00000073,
	"pageUp":          0x00000074,
	"delete":          0x00000075,
	"end":             0x00000077,
	"pageDown":        0x00000079,
	"arrowRight":      0x0000007c,
	"arrowLeft":       0x0000007b,
	"arrowDown":       0x0000007d,
	"arrowUp":         0x0000007e,
	"numLock":         0x00000047,
	"numpadDivide":    0x0000004b,
	"numpadMultiply":  0x00000043,
	"numpadSubtract":  0x0000004e,
	"numpadAdd":       0x00000045,
	"numpadEnter":     0x0000004c,
	"numpad1":         0x00000053,
	"numpad2":         0x00000054,
	"numpad3":         0x00000055,
	"numpad4":         0x00000056,
	"numpad5":         0x00000057,
	"numpad6":         0x00000058,
	"numpad7":         0x00000059,
	"numpad8":         0x0000005b,
	"numpad9":         0x0000005c,
	"numpad0":         0x00000052,
	"numpadDecimal":   0x00000041,
	"intlBackslash":   0x0000000a,
	"contextMenu":     0x0000006e,
	"numpadEqual":     0x00000051,
	"f13":             0x00000069,
	"f14":             0x0000006b,
	"f15":             0x00000071,
	"f16":             0x0000006a,
	"f17":             0x00000040,
	"f18":             0x0000004f,
	"f19":             0x00000050,
	"f20":             0x0000005a,
	"audioVolumeMute": 0x0000004a,
	"audioVolumeUp":   0x00000048,
	"audioVolumeDown": 0x00000049,
	"numpadComma":     0x0000005f,
	"intlRo":          0x0000005e,
	"intlYen":         0x0000005d,
	"lang1":           0x00000068,
	"lang2":           0x00000066,
	"controlLeft":     0x0000003b,
	"shiftLeft":       0x00000038,
	"altLeft":         0x0000003a,
	"metaLeft":        0x00000037,
	"controlRight":    0x0000003e,
	"shiftRight":      0x0000003c,
	"altRight":        0x0000003d,
	"metaRight":       0x00000036,
	"fn":              0x0000003f,
}