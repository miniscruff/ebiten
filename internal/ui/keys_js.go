// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build js

package ui

var keyCodeToKey = map[int]Key{
	8:   KeyBackspace,
	9:   KeyTab,
	13:  KeyEnter,
	16:  KeyShift,
	17:  KeyControl,
	18:  KeyAlt,
	20:  KeyCapsLock,
	27:  KeyEscape,
	32:  KeySpace,
	33:  KeyPageUp,
	34:  KeyPageDown,
	35:  KeyEnd,
	36:  KeyHome,
	37:  KeyLeft,
	38:  KeyUp,
	39:  KeyRight,
	40:  KeyDown,
	45:  KeyInsert,
	46:  KeyDelete,
	48:  Key0,
	49:  Key1,
	50:  Key2,
	51:  Key3,
	52:  Key4,
	53:  Key5,
	54:  Key6,
	55:  Key7,
	56:  Key8,
	57:  Key9,
	65:  KeyA,
	66:  KeyB,
	67:  KeyC,
	68:  KeyD,
	69:  KeyE,
	70:  KeyF,
	71:  KeyG,
	72:  KeyH,
	73:  KeyI,
	74:  KeyJ,
	75:  KeyK,
	76:  KeyL,
	77:  KeyM,
	78:  KeyN,
	79:  KeyO,
	80:  KeyP,
	81:  KeyQ,
	82:  KeyR,
	83:  KeyS,
	84:  KeyT,
	85:  KeyU,
	86:  KeyV,
	87:  KeyW,
	88:  KeyX,
	89:  KeyY,
	90:  KeyZ,
	112: KeyF1,
	113: KeyF2,
	114: KeyF3,
	115: KeyF4,
	116: KeyF5,
	117: KeyF6,
	118: KeyF7,
	119: KeyF8,
	120: KeyF9,
	121: KeyF10,
	122: KeyF11,
	123: KeyF12,
	188: KeyComma,
	190: KeyPeriod,
}
