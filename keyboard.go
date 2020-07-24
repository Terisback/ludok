package ludok

import "github.com/veandco/go-sdl2/sdl"

type KeyboardEvent struct {
	// Event type
	//
	// Pressed, Released
	State InputEventType
	// Physical key code
	//
	// Used to represent the physical location of a keyboard key on the keyboard
	Scancode uint
	// Virtual key code
	//
	// Used to represent keyboard keys using the current layout of the keyboard
	Symcode KeyboardKeycode
	// Keyboard event modifier
	//
	// May have more than one modifier, better to check with Mod.Has()
	Mod KeyboardEventModifier
	// Is repeated, non-zero if yes
	//
	// It's very laggy, please use GetKeyboardState() if you want to get something instantly
	Repeated uint8
}

type InputEventType int

const (
	Pressed  InputEventType = sdl.PRESSED
	Released InputEventType = sdl.RELEASED
)

type KeyboardKeycode uint32

const (
	K_UNKNOWN KeyboardKeycode = sdl.K_UNKNOWN // "" (no name, empty string)

	K_RETURN     KeyboardKeycode = sdl.K_RETURN     // "Return" (the Enter key (main keyboard))
	K_ESCAPE     KeyboardKeycode = sdl.K_ESCAPE     // "Escape" (the Esc key)
	K_BACKSPACE  KeyboardKeycode = sdl.K_BACKSPACE  // "Backspace"
	K_TAB        KeyboardKeycode = sdl.K_TAB        // "Tab" (the Tab key)
	K_SPACE      KeyboardKeycode = sdl.K_SPACE      // "Space" (the Space Bar key(s))
	K_EXCLAIM    KeyboardKeycode = sdl.K_EXCLAIM    // "!"
	K_QUOTEDBL   KeyboardKeycode = sdl.K_QUOTEDBL   // """
	K_HASH       KeyboardKeycode = sdl.K_HASH       // "#"
	K_PERCENT    KeyboardKeycode = sdl.K_PERCENT    // "%"
	K_DOLLAR     KeyboardKeycode = sdl.K_DOLLAR     // "$"
	K_AMPERSAND  KeyboardKeycode = sdl.K_AMPERSAND  // "&"
	K_QUOTE      KeyboardKeycode = sdl.K_QUOTE      // "'"
	K_LEFTPAREN  KeyboardKeycode = sdl.K_LEFTPAREN  // "("
	K_RIGHTPAREN KeyboardKeycode = sdl.K_RIGHTPAREN // ")"
	K_ASTERISK   KeyboardKeycode = sdl.K_ASTERISK   // "*"
	K_PLUS       KeyboardKeycode = sdl.K_PLUS       // "+"
	K_COMMA      KeyboardKeycode = sdl.K_COMMA      // ","
	K_MINUS      KeyboardKeycode = sdl.K_MINUS      // "-"
	K_PERIOD     KeyboardKeycode = sdl.K_PERIOD     // "."
	K_SLASH      KeyboardKeycode = sdl.K_SLASH      // "/"
	K_0          KeyboardKeycode = sdl.K_0          // "0"
	K_1          KeyboardKeycode = sdl.K_1          // "1"
	K_2          KeyboardKeycode = sdl.K_2          // "2"
	K_3          KeyboardKeycode = sdl.K_3          // "3"
	K_4          KeyboardKeycode = sdl.K_4          // "4"
	K_5          KeyboardKeycode = sdl.K_5          // "5"
	K_6          KeyboardKeycode = sdl.K_6          // "6"
	K_7          KeyboardKeycode = sdl.K_7          // "7"
	K_8          KeyboardKeycode = sdl.K_8          // "8"
	K_9          KeyboardKeycode = sdl.K_9          // "9"
	K_COLON      KeyboardKeycode = sdl.K_COLON      // ":"
	K_SEMICOLON  KeyboardKeycode = sdl.K_SEMICOLON  // ";"
	K_LESS       KeyboardKeycode = sdl.K_LESS       // "<"
	K_EQUALS     KeyboardKeycode = sdl.K_EQUALS     // "="
	K_GREATER    KeyboardKeycode = sdl.K_GREATER    // ">"
	K_QUESTION   KeyboardKeycode = sdl.K_QUESTION   // "?"
	K_AT         KeyboardKeycode = sdl.K_AT         // "@"
	/*
	   Skip uppercase letters
	*/
	K_LEFTBRACKET  KeyboardKeycode = sdl.K_LEFTBRACKET  // "["
	K_BACKSLASH    KeyboardKeycode = sdl.K_BACKSLASH    // "\"
	K_RIGHTBRACKET KeyboardKeycode = sdl.K_RIGHTBRACKET // "]"
	K_CARET        KeyboardKeycode = sdl.K_CARET        // "^"
	K_UNDERSCORE   KeyboardKeycode = sdl.K_UNDERSCORE   // "_"
	K_BACKQUOTE    KeyboardKeycode = sdl.K_BACKQUOTE    // "`"
	K_a            KeyboardKeycode = sdl.K_a            // "A"
	K_b            KeyboardKeycode = sdl.K_b            // "B"
	K_c            KeyboardKeycode = sdl.K_c            // "C"
	K_d            KeyboardKeycode = sdl.K_d            // "D"
	K_e            KeyboardKeycode = sdl.K_e            // "E"
	K_f            KeyboardKeycode = sdl.K_f            // "F"
	K_g            KeyboardKeycode = sdl.K_g            // "G"
	K_h            KeyboardKeycode = sdl.K_h            // "H"
	K_i            KeyboardKeycode = sdl.K_i            // "I"
	K_j            KeyboardKeycode = sdl.K_j            // "J"
	K_k            KeyboardKeycode = sdl.K_k            // "K"
	K_l            KeyboardKeycode = sdl.K_l            // "L"
	K_m            KeyboardKeycode = sdl.K_m            // "M"
	K_n            KeyboardKeycode = sdl.K_n            // "N"
	K_o            KeyboardKeycode = sdl.K_o            // "O"
	K_p            KeyboardKeycode = sdl.K_p            // "P"
	K_q            KeyboardKeycode = sdl.K_q            // "Q"
	K_r            KeyboardKeycode = sdl.K_r            // "R"
	K_s            KeyboardKeycode = sdl.K_s            // "S"
	K_t            KeyboardKeycode = sdl.K_t            // "T"
	K_u            KeyboardKeycode = sdl.K_u            // "U"
	K_v            KeyboardKeycode = sdl.K_v            // "V"
	K_w            KeyboardKeycode = sdl.K_w            // "W"
	K_x            KeyboardKeycode = sdl.K_x            // "X"
	K_y            KeyboardKeycode = sdl.K_y            // "Y"
	K_z            KeyboardKeycode = sdl.K_z            // "Z"

	K_CAPSLOCK = sdl.K_CAPSLOCK // "CapsLock"

	K_F1  KeyboardKeycode = sdl.K_F1  // "F1"
	K_F2  KeyboardKeycode = sdl.K_F2  // "F2"
	K_F3  KeyboardKeycode = sdl.K_F3  // "F3"
	K_F4  KeyboardKeycode = sdl.K_F4  // "F4"
	K_F5  KeyboardKeycode = sdl.K_F5  // "F5"
	K_F6  KeyboardKeycode = sdl.K_F6  // "F6"
	K_F7  KeyboardKeycode = sdl.K_F7  // "F7"
	K_F8  KeyboardKeycode = sdl.K_F8  // "F8"
	K_F9  KeyboardKeycode = sdl.K_F9  // "F9"
	K_F10 KeyboardKeycode = sdl.K_F10 // "F10"
	K_F11 KeyboardKeycode = sdl.K_F11 // "F11"
	K_F12 KeyboardKeycode = sdl.K_F12 // "F12"

	K_PRINTSCREEN KeyboardKeycode = sdl.K_PRINTSCREEN // "PrintScreen"
	K_SCROLLLOCK  KeyboardKeycode = sdl.K_SCROLLLOCK  // "ScrollLock"
	K_PAUSE       KeyboardKeycode = sdl.K_PAUSE       // "Pause" (the Pause / Break key)
	K_INSERT      KeyboardKeycode = sdl.K_INSERT      // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	K_HOME        KeyboardKeycode = sdl.K_HOME        // "Home"
	K_PAGEUP      KeyboardKeycode = sdl.K_PAGEUP      // "PageUp"
	K_DELETE      KeyboardKeycode = sdl.K_DELETE      // "Delete"
	K_END         KeyboardKeycode = sdl.K_END         // "End"
	K_PAGEDOWN    KeyboardKeycode = sdl.K_PAGEDOWN    // "PageDown"
	K_RIGHT       KeyboardKeycode = sdl.K_RIGHT       // "Right" (the Right arrow key (navigation keypad))
	K_LEFT        KeyboardKeycode = sdl.K_LEFT        // "Left" (the Left arrow key (navigation keypad))
	K_DOWN        KeyboardKeycode = sdl.K_DOWN        // "Down" (the Down arrow key (navigation keypad))
	K_UP          KeyboardKeycode = sdl.K_UP          // "Up" (the Up arrow key (navigation keypad))

	K_NUMLOCKCLEAR KeyboardKeycode = sdl.K_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	K_KP_DIVIDE    KeyboardKeycode = sdl.K_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	K_KP_MULTIPLY  KeyboardKeycode = sdl.K_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	K_KP_MINUS     KeyboardKeycode = sdl.K_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	K_KP_PLUS      KeyboardKeycode = sdl.K_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	K_KP_ENTER     KeyboardKeycode = sdl.K_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	K_KP_1         KeyboardKeycode = sdl.K_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	K_KP_2         KeyboardKeycode = sdl.K_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	K_KP_3         KeyboardKeycode = sdl.K_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	K_KP_4         KeyboardKeycode = sdl.K_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	K_KP_5         KeyboardKeycode = sdl.K_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	K_KP_6         KeyboardKeycode = sdl.K_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	K_KP_7         KeyboardKeycode = sdl.K_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	K_KP_8         KeyboardKeycode = sdl.K_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	K_KP_9         KeyboardKeycode = sdl.K_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	K_KP_0         KeyboardKeycode = sdl.K_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	K_KP_PERIOD    KeyboardKeycode = sdl.K_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	K_APPLICATION    KeyboardKeycode = sdl.K_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	K_POWER          KeyboardKeycode = sdl.K_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	K_KP_EQUALS      KeyboardKeycode = sdl.K_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	K_F13            KeyboardKeycode = sdl.K_F13            // "F13"
	K_F14            KeyboardKeycode = sdl.K_F14            // "F14"
	K_F15            KeyboardKeycode = sdl.K_F15            // "F15"
	K_F16            KeyboardKeycode = sdl.K_F16            // "F16"
	K_F17            KeyboardKeycode = sdl.K_F17            // "F17"
	K_F18            KeyboardKeycode = sdl.K_F18            // "F18"
	K_F19            KeyboardKeycode = sdl.K_F19            // "F19"
	K_F20            KeyboardKeycode = sdl.K_F20            // "F20"
	K_F21            KeyboardKeycode = sdl.K_F21            // "F21"
	K_F22            KeyboardKeycode = sdl.K_F22            // "F22"
	K_F23            KeyboardKeycode = sdl.K_F23            // "F23"
	K_F24            KeyboardKeycode = sdl.K_F24            // "F24"
	K_EXECUTE        KeyboardKeycode = sdl.K_EXECUTE        // "Execute"
	K_HELP           KeyboardKeycode = sdl.K_HELP           // "Help"
	K_MENU           KeyboardKeycode = sdl.K_MENU           // "Menu"
	K_SELECT         KeyboardKeycode = sdl.K_SELECT         // "Select"
	K_STOP           KeyboardKeycode = sdl.K_STOP           // "Stop"
	K_AGAIN          KeyboardKeycode = sdl.K_AGAIN          // "Again" (the Again key (Redo))
	K_UNDO           KeyboardKeycode = sdl.K_UNDO           // "Undo"
	K_CUT            KeyboardKeycode = sdl.K_CUT            // "Cut"
	K_COPY           KeyboardKeycode = sdl.K_COPY           // "Copy"
	K_PASTE          KeyboardKeycode = sdl.K_PASTE          // "Paste"
	K_FIND           KeyboardKeycode = sdl.K_FIND           // "Find"
	K_MUTE           KeyboardKeycode = sdl.K_MUTE           // "Mute"
	K_VOLUMEUP       KeyboardKeycode = sdl.K_VOLUMEUP       // "VolumeUp"
	K_VOLUMEDOWN     KeyboardKeycode = sdl.K_VOLUMEDOWN     // "VolumeDown"
	K_KP_COMMA       KeyboardKeycode = sdl.K_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	K_KP_EQUALSAS400 KeyboardKeycode = sdl.K_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	K_ALTERASE   KeyboardKeycode = sdl.K_ALTERASE   // "AltErase" (Erase-Eaze)
	K_SYSREQ     KeyboardKeycode = sdl.K_SYSREQ     // "SysReq" (the SysReq key)
	K_CANCEL     KeyboardKeycode = sdl.K_CANCEL     // "Cancel"
	K_CLEAR      KeyboardKeycode = sdl.K_CLEAR      // "Clear"
	K_PRIOR      KeyboardKeycode = sdl.K_PRIOR      // "Prior"
	K_RETURN2    KeyboardKeycode = sdl.K_RETURN2    // "Return"
	K_SEPARATOR  KeyboardKeycode = sdl.K_SEPARATOR  // "Separator"
	K_OUT        KeyboardKeycode = sdl.K_OUT        // "Out"
	K_OPER       KeyboardKeycode = sdl.K_OPER       // "Oper"
	K_CLEARAGAIN KeyboardKeycode = sdl.K_CLEARAGAIN // "Clear / Again"
	K_CRSEL      KeyboardKeycode = sdl.K_CRSEL      // "CrSel"
	K_EXSEL      KeyboardKeycode = sdl.K_EXSEL      // "ExSel"

	K_KP_00              KeyboardKeycode = sdl.K_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	K_KP_000             KeyboardKeycode = sdl.K_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	K_THOUSANDSSEPARATOR KeyboardKeycode = sdl.K_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	K_DECIMALSEPARATOR   KeyboardKeycode = sdl.K_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	K_CURRENCYUNIT       KeyboardKeycode = sdl.K_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	K_CURRENCYSUBUNIT    KeyboardKeycode = sdl.K_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	K_KP_LEFTPAREN       KeyboardKeycode = sdl.K_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	K_KP_RIGHTPAREN      KeyboardKeycode = sdl.K_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	K_KP_LEFTBRACE       KeyboardKeycode = sdl.K_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	K_KP_RIGHTBRACE      KeyboardKeycode = sdl.K_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	K_KP_TAB             KeyboardKeycode = sdl.K_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	K_KP_BACKSPACE       KeyboardKeycode = sdl.K_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	K_KP_A               KeyboardKeycode = sdl.K_KP_A               // "Keypad A" (the A key (numeric keypad))
	K_KP_B               KeyboardKeycode = sdl.K_KP_B               // "Keypad B" (the B key (numeric keypad))
	K_KP_C               KeyboardKeycode = sdl.K_KP_C               // "Keypad C" (the C key (numeric keypad))
	K_KP_D               KeyboardKeycode = sdl.K_KP_D               // "Keypad D" (the D key (numeric keypad))
	K_KP_E               KeyboardKeycode = sdl.K_KP_E               // "Keypad E" (the E key (numeric keypad))
	K_KP_F               KeyboardKeycode = sdl.K_KP_F               // "Keypad F" (the F key (numeric keypad))
	K_KP_XOR             KeyboardKeycode = sdl.K_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	K_KP_POWER           KeyboardKeycode = sdl.K_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	K_KP_PERCENT         KeyboardKeycode = sdl.K_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	K_KP_LESS            KeyboardKeycode = sdl.K_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	K_KP_GREATER         KeyboardKeycode = sdl.K_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	K_KP_AMPERSAND       KeyboardKeycode = sdl.K_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	K_KP_DBLAMPERSAND    KeyboardKeycode = sdl.K_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	K_KP_VERTICALBAR     KeyboardKeycode = sdl.K_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	K_KP_DBLVERTICALBAR  KeyboardKeycode = sdl.K_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	K_KP_COLON           KeyboardKeycode = sdl.K_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	K_KP_HASH            KeyboardKeycode = sdl.K_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	K_KP_SPACE           KeyboardKeycode = sdl.K_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	K_KP_AT              KeyboardKeycode = sdl.K_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	K_KP_EXCLAM          KeyboardKeycode = sdl.K_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	K_KP_MEMSTORE        KeyboardKeycode = sdl.K_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	K_KP_MEMRECALL       KeyboardKeycode = sdl.K_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	K_KP_MEMCLEAR        KeyboardKeycode = sdl.K_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	K_KP_MEMADD          KeyboardKeycode = sdl.K_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	K_KP_MEMSUBTRACT     KeyboardKeycode = sdl.K_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	K_KP_MEMMULTIPLY     KeyboardKeycode = sdl.K_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	K_KP_MEMDIVIDE       KeyboardKeycode = sdl.K_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	K_KP_PLUSMINUS       KeyboardKeycode = sdl.K_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	K_KP_CLEAR           KeyboardKeycode = sdl.K_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	K_KP_CLEARENTRY      KeyboardKeycode = sdl.K_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	K_KP_BINARY          KeyboardKeycode = sdl.K_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	K_KP_OCTAL           KeyboardKeycode = sdl.K_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	K_KP_DECIMAL         KeyboardKeycode = sdl.K_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	K_KP_HEXADECIMAL     KeyboardKeycode = sdl.K_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	K_LCTRL  KeyboardKeycode = sdl.K_LCTRL  // "Left Ctrl"
	K_LSHIFT KeyboardKeycode = sdl.K_LSHIFT // "Left Shift"
	K_LALT   KeyboardKeycode = sdl.K_LALT   // "Left Alt" (alt, option)
	K_LGUI   KeyboardKeycode = sdl.K_LGUI   // "Left GUI" (windows, command (apple), meta)
	K_RCTRL  KeyboardKeycode = sdl.K_RCTRL  // "Right Ctrl"
	K_RSHIFT KeyboardKeycode = sdl.K_RSHIFT // "Right Shift"
	K_RALT   KeyboardKeycode = sdl.K_RALT   // "Right Alt" (alt, option)
	K_RGUI   KeyboardKeycode = sdl.K_RGUI   // "Right GUI" (windows, command (apple), meta)

	K_MODE KeyboardKeycode = sdl.K_MODE // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)

	K_AUDIONEXT    KeyboardKeycode = sdl.K_AUDIONEXT    // "AudioNext" (the Next Track media key)
	K_AUDIOPREV    KeyboardKeycode = sdl.K_AUDIOPREV    // "AudioPrev" (the Previous Track media key)
	K_AUDIOSTOP    KeyboardKeycode = sdl.K_AUDIOSTOP    // "AudioStop" (the Stop media key)
	K_AUDIOPLAY    KeyboardKeycode = sdl.K_AUDIOPLAY    // "AudioPlay" (the Play media key)
	K_AUDIOMUTE    KeyboardKeycode = sdl.K_AUDIOMUTE    // "AudioMute" (the Mute volume key)
	K_MEDIASELECT  KeyboardKeycode = sdl.K_MEDIASELECT  // "MediaSelect" (the Media Select key)
	K_WWW          KeyboardKeycode = sdl.K_WWW          // "WWW" (the WWW/World Wide Web key)
	K_MAIL         KeyboardKeycode = sdl.K_MAIL         // "Mail" (the Mail/eMail key)
	K_CALCULATOR   KeyboardKeycode = sdl.K_CALCULATOR   // "Calculator" (the Calculator key)
	K_COMPUTER     KeyboardKeycode = sdl.K_COMPUTER     // "Computer" (the My Computer key)
	K_AC_SEARCH    KeyboardKeycode = sdl.K_AC_SEARCH    // "AC Search" (the Search key (application control keypad))
	K_AC_HOME      KeyboardKeycode = sdl.K_AC_HOME      // "AC Home" (the Home key (application control keypad))
	K_AC_BACK      KeyboardKeycode = sdl.K_AC_BACK      // "AC Back" (the Back key (application control keypad))
	K_AC_FORWARD   KeyboardKeycode = sdl.K_AC_FORWARD   // "AC Forward" (the Forward key (application control keypad))
	K_AC_STOP      KeyboardKeycode = sdl.K_AC_STOP      // "AC Stop" (the Stop key (application control keypad))
	K_AC_REFRESH   KeyboardKeycode = sdl.K_AC_REFRESH   // "AC Refresh" (the Refresh key (application control keypad))
	K_AC_BOOKMARKS KeyboardKeycode = sdl.K_AC_BOOKMARKS // "AC Bookmarks" (the Bookmarks key (application control keypad))

	K_BRIGHTNESSDOWN KeyboardKeycode = sdl.K_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	K_BRIGHTNESSUP   KeyboardKeycode = sdl.K_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	K_DISPLAYSWITCH  KeyboardKeycode = sdl.K_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	K_KBDILLUMTOGGLE KeyboardKeycode = sdl.K_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	K_KBDILLUMDOWN   KeyboardKeycode = sdl.K_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	K_KBDILLUMUP     KeyboardKeycode = sdl.K_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	K_EJECT          KeyboardKeycode = sdl.K_EJECT          // "Eject" (the Eject key)
	K_SLEEP          KeyboardKeycode = sdl.K_SLEEP          // "Sleep" (the Sleep key)
)

type KeyboardEventModifier uint16

const (
	KMOD_LSHIFT KeyboardEventModifier = sdl.KMOD_LSHIFT // the left Shift key is down
	KMOD_RSHIFT KeyboardEventModifier = sdl.KMOD_RSHIFT // the right Shift key is down
	KMOD_LCTRL  KeyboardEventModifier = sdl.KMOD_LCTRL  // the left Ctrl (Control) key is down
	KMOD_RCTRL  KeyboardEventModifier = sdl.KMOD_RCTRL  // the right Ctrl (Control) key is down
	KMOD_LALT   KeyboardEventModifier = sdl.KMOD_LALT   // the left Alt key is down
	KMOD_RALT   KeyboardEventModifier = sdl.KMOD_RALT   // the right Alt key is down
	KMOD_LGUI   KeyboardEventModifier = sdl.KMOD_LGUI   // the left GUI key (often the Windows key) is down
	KMOD_RGUI   KeyboardEventModifier = sdl.KMOD_RGUI   // the right GUI key (often the Windows key) is down
	KMOD_NUM    KeyboardEventModifier = sdl.KMOD_NUM    // the Num Lock key (may be located on an extended keypad) is down
	KMOD_CAPS   KeyboardEventModifier = sdl.KMOD_CAPS   // the Caps Lock key is down
	KMOD_MODE   KeyboardEventModifier = sdl.KMOD_MODE   // the AltGr key is down
	KMOD_CTRL   KeyboardEventModifier = sdl.KMOD_CTRL   // (KMOD_LCTRL|KMOD_RCTRL)
	KMOD_SHIFT  KeyboardEventModifier = sdl.KMOD_SHIFT  // (KMOD_LSHIFT|KMOD_RSHIFT)
	KMOD_ALT    KeyboardEventModifier = sdl.KMOD_ALT    // (KMOD_LALT|KMOD_RALT)
	KMOD_GUI    KeyboardEventModifier = sdl.KMOD_GUI    // (KMOD_LGUI|KMOD_RGUI)
)

func (k KeyboardEventModifier) Has(mod KeyboardEventModifier) bool {
	return k&mod != 0
}
