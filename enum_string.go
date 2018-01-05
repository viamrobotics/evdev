package evdev

import "strconv"

// Code generated by gen.sh; DO NOT EDIT.

const (
	_AbsoluteType_name_0 = "XYZRXRYRZThrottleRudderWheelGasBrake"
	_AbsoluteType_name_1 = "Hat0XHat0YHat1XHat1YHat2XHat2YHat3XHat3YPressureDistanceTiltXTiltYToolWidth"
	_AbsoluteType_name_2 = "Volume"
	_AbsoluteType_name_3 = "Misc"
	_AbsoluteType_name_4 = "MTSlotMTTouchMajorMTTouchMinorMTWidthMajorMTWidthMinorMTOrientationMTPositionXMTPositionYMTToolTYPEMTBlobIdMTTrackingIdMTPressureMTDistanceMTToolXMTToolY"
)

var (
	_AbsoluteType_index_0 = [...]uint8{0, 1, 2, 3, 5, 7, 9, 17, 23, 28, 31, 36}
	_AbsoluteType_index_1 = [...]uint8{0, 5, 10, 15, 20, 25, 30, 35, 40, 48, 56, 61, 66, 75}
	_AbsoluteType_index_4 = [...]uint8{0, 6, 18, 30, 42, 54, 67, 78, 89, 99, 107, 119, 129, 139, 146, 153}
)

func (i AbsoluteType) String() string {
	switch {
	case 0 <= i && i <= 10:
		return _AbsoluteType_name_0[_AbsoluteType_index_0[i]:_AbsoluteType_index_0[i+1]]
	case 16 <= i && i <= 28:
		i -= 16
		return _AbsoluteType_name_1[_AbsoluteType_index_1[i]:_AbsoluteType_index_1[i+1]]
	case i == 32:
		return _AbsoluteType_name_2
	case i == 40:
		return _AbsoluteType_name_3
	case 47 <= i && i <= 61:
		i -= 47
		return _AbsoluteType_name_4[_AbsoluteType_index_4[i]:_AbsoluteType_index_4[i+1]]
	default:
		return "AbsoluteType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const (
	_BusType_name_0 = "PCIISAPNPUSBHILBluetoothVirtual"
	_BusType_name_1 = "ISAI8042XTKBDRS232GamePortParallelPortAmigaADBI2CHostGSCAtariSPIRMICECIntelISHTP"
)

var (
	_BusType_index_0 = [...]uint8{0, 3, 9, 12, 15, 24, 31}
	_BusType_index_1 = [...]uint8{0, 3, 8, 13, 18, 26, 38, 43, 46, 49, 53, 56, 61, 64, 67, 70, 80}
)

func (i BusType) String() string {
	switch {
	case 1 <= i && i <= 6:
		i -= 1
		return _BusType_name_0[_BusType_index_0[i]:_BusType_index_0[i+1]]
	case 16 <= i && i <= 31:
		i -= 16
		return _BusType_name_1[_BusType_index_1[i]:_BusType_index_1[i+1]]
	default:
		return "BusType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _EffectType_name = "RumblePeriodicConstantSpringFrictionDamperInertiaRampSquareTriangleSineSawUpSawDownCustom"

var _EffectType_index = [...]uint8{0, 6, 14, 22, 28, 36, 42, 49, 53, 59, 67, 71, 76, 83, 89}

func (i EffectType) String() string {
	i -= 80
	if i < 0 || i >= EffectType(len(_EffectType_index)-1) {
		return "EffectType(" + strconv.FormatInt(int64(i+80), 10) + ")"
	}
	return _EffectType_name[_EffectType_index[i]:_EffectType_index[i+1]]
}

const (
	_EffectDirType_name_0 = "Down"
	_EffectDirType_name_1 = "Left"
	_EffectDirType_name_2 = "Up"
	_EffectDirType_name_3 = "Right"
)

func (i EffectDirType) String() string {
	switch {
	case i == 0:
		return _EffectDirType_name_0
	case i == 16384:
		return _EffectDirType_name_1
	case i == 32768:
		return _EffectDirType_name_2
	case i == 49152:
		return _EffectDirType_name_3
	default:
		return "EffectDirType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _EffectPropType_name = "GainAutoCenter"

var _EffectPropType_index = [...]uint8{0, 4, 14}

func (i EffectPropType) String() string {
	i -= 96
	if i < 0 || i >= EffectPropType(len(_EffectPropType_index)-1) {
		return "EffectPropType(" + strconv.FormatInt(int64(i+96), 10) + ")"
	}
	return _EffectPropType_name[_EffectPropType_index[i]:_EffectPropType_index[i+1]]
}

const _EffectStatusType_name = "StoppedPlaying"

var _EffectStatusType_index = [...]uint8{0, 7, 14}

func (i EffectStatusType) String() string {
	if i < 0 || i >= EffectStatusType(len(_EffectStatusType_index)-1) {
		return "EffectStatusType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EffectStatusType_name[_EffectStatusType_index[i]:_EffectStatusType_index[i+1]]
}

const (
	_EventType_name_0 = "SyncKeyRelativeAbsoluteMiscSwitch"
	_EventType_name_1 = "LEDSound"
	_EventType_name_2 = "RepeatEffectPowerEffectStatus"
)

var (
	_EventType_index_0 = [...]uint8{0, 4, 7, 15, 23, 27, 33}
	_EventType_index_1 = [...]uint8{0, 3, 8}
	_EventType_index_2 = [...]uint8{0, 6, 12, 17, 29}
)

func (i EventType) String() string {
	switch {
	case 0 <= i && i <= 5:
		return _EventType_name_0[_EventType_index_0[i]:_EventType_index_0[i+1]]
	case 17 <= i && i <= 18:
		i -= 17
		return _EventType_name_1[_EventType_index_1[i]:_EventType_index_1[i+1]]
	case 20 <= i && i <= 23:
		i -= 20
		return _EventType_name_2[_EventType_index_2[i]:_EventType_index_2[i+1]]
	default:
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _InputType_name = "KeyboardMouseJoystick"

var _InputType_index = [...]uint8{0, 8, 13, 21}

func (i InputType) String() string {
	if i < 0 || i >= InputType(len(_InputType_index)-1) {
		return "InputType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InputType_name[_InputType_index[i]:_InputType_index[i+1]]
}

const _InputPropType_name = "PointerDirectButtonPadSemiMT"

var _InputPropType_index = [...]uint8{0, 7, 13, 22, 28}

func (i InputPropType) String() string {
	if i < 0 || i >= InputPropType(len(_InputPropType_index)-1) {
		return "InputPropType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InputPropType_name[_InputPropType_index[i]:_InputPropType_index[i+1]]
}

const _KeyType_name = "ReservedEscape1234567890MinusEqualBackSpaceTabQWERTYUIOPLeftBraceRightBraceEnterLeftCtrlASDFGHJKLSemiColonApostropheGraveLeftShiftBackSlashZXCVBNMCommaDotSlashRightShiftKPAsteriskLeftAltSpaceCapsLockF1F2F3F4F5F6F7F8F9F10NumLockScrollLockKP7KP8KP9KPMinusKP4KP5KP6KPPlusKP1KP2KP3KP0KPDotZenkakuhankaku102NDF11F12ROKatakanaHiraganaHenkanKatakanaHiraganaMuhenkanKPJPCommaKPEnterRightCtrlKPSlashSysRQRightAltLineFeedHomeUpPageUpLeftRightEndDownPageDownInsertDeleteMacroMuteVolumeDownVolumeUpPowerKPEqualKPPlusMinusPauseScaleKPCommaHangulHanjaYenLeftMetaRightMetaComposeStopAgainPropsUndoFrontCopyOpenPasteFindCutHelpMenuCalcSetupSleepWakeupFileSendFileDeleteFileXFerProg1Prog2WWWMSDOSScreenlockDirectionCycleWindowsMailBookmarksComputerBackForwardCloseCDEjectCDEjectCloseCDNextSongPlayPausePreviousSongStopCDRecordRewindPhoneISOConfigHomepageRefreshExitMoveEditScrollUpScrollDownKPLeftParenKPRightParenNewRedoF13F14F15F16F17F18F19F20F21F22F23F24PlayCDPauseCDProg3Prog4DashboardSuspendClosePlayFastForwardBassBoostPrintHPCaneraSoundQuestionEmailChatSearchConnectFinanceSportShopAltEraseCancelBrightnessDownBrightnessUpMediaSwitchVideoModeKBDIllumToggleKBDIllumDownKBDIllumUpSendReplyForwardMailSaveDocumentsBatteryBluetoothWLANUWBUnknownVideoNextVideoPreviousBrightnessCycleBrightnessZeroDisplayOffWIMaxRFKillMicMuteBtn0Btn1Btn2Btn3Btn4Btn5Btn6Btn7Btn8Btn9BtnMouseBtnRightBtnMiddleBtnSideBtnExtraBtnForwardBtnBackBtnTaskBtnJoystickBtnThumbBtnThumb2BtnTopBtnTop2BtnPinkieBtnBaseBtnBase2BtnBase3BtnBase4BtnBase5BtnBase6BtnDeadBtnGamepadBtnBBtnCBtnXBtnYBtnZBtnTLBtnTRBtnTL2BtnTR2BtnSelectBtnStartBtnModeBtnThumbLBtnThumbRBtnDigiBtnTooLRubberBtnToolBrushBtnToolPencilBtnToolAirbrushBtnToolFingerBtnToolMouseBtnToolLensBtnToolQuintTapBtnTouchBtnStylusBtnStylus2BtnToolDoubleTapBtnToolTrippleTapBtnToolQuadTapBtnWheelBtnGearUpOkSelectGotoClearPower2OptionInfoTimeVendorArchiveProgramChannelFavoritesEPGPVRMHPLanguageTitleSubtitleAngleZoomModeKeyboardScreenPCTVTV2VCRVCR2SATSAT2CDTapeRadioTunerPlayerTextDVDAUXMP3AudioVideoDirectoryListMemoCalenderRedGreenYellowBlueChannelUpChannelDownFirstLastABNextRestartSlowShuffleBreakPreviousDigitsTeenTwenVideoPhoneGamesZoomInZoomOutZoomResetWordProcessorEditorSpreadsheetGraphicsEditorPresentationDatabaseNewsVoiceMailAddressBookMessengerDisplayToggleSpellCheckLogoffDollarEuroFrameBackframeForwardContextMenuMediaRepeat10ChannelsUp10ChannelsDownImagesDelEOLDelEOSInsLineDelLineFNFNEscFNF1FNF2FNF3FNF4FNF5FNF6FNF7FNF8FNF9FNF10FNF11FNF12FN1FN2FNDFNEFNFFNSFNBBRLDot1BRLDot2BRLDot3BRLDot4BRLDot5BRLDot6BRLDot7BRLDot8BRLDot9BRLDot10Numeric0Numeric1Numeric2Numeric3Numeric4Numeric5Numeric6Numeric7Numeric8Numeric9NumericStarNumericPoundCameraFocusWPSButtonTouchpadToggleTouchpadOnTouchpadOffCameraZoomInCameraZoomOutCameraUpCameraDownCameraLeftCameraRightAttendantOnAttendantOffAttendantToggleLightsToggleBtnTriggerHappyBtnTriggerHappy2BtnTriggerHappy3BtnTriggerHappy4BtnTriggerHappy5BtnTriggerHappy6BtnTriggerHappy7BtnTriggerHappy8BtnTriggerHappy9BtnTriggerHappy10BtnTriggerHappy11BtnTriggerHappy12BtnTriggerHappy13BtnTriggerHappy14BtnTriggerHappy15BtnTriggerHappy16BtnTriggerHappy17BtnTriggerHappy18BtnTriggerHappy19BtnTriggerHappy20BtnTriggerHappy21BtnTriggerHappy22BtnTriggerHappy23BtnTriggerHappy24BtnTriggerHappy25BtnTriggerHappy26BtnTriggerHappy27BtnTriggerHappy28BtnTriggerHappy29BtnTriggerHappy30BtnTriggerHappy31BtnTriggerHappy32BtnTriggerHappy33BtnTriggerHappy34BtnTriggerHappy35BtnTriggerHappy36BtnTriggerHappy37BtnTriggerHappy38BtnTriggerHappy39BtnTriggerHappy40"

var _KeyType_map = map[KeyType]string{
	0:   _KeyType_name[0:8],
	1:   _KeyType_name[8:14],
	2:   _KeyType_name[14:15],
	3:   _KeyType_name[15:16],
	4:   _KeyType_name[16:17],
	5:   _KeyType_name[17:18],
	6:   _KeyType_name[18:19],
	7:   _KeyType_name[19:20],
	8:   _KeyType_name[20:21],
	9:   _KeyType_name[21:22],
	10:  _KeyType_name[22:23],
	11:  _KeyType_name[23:24],
	12:  _KeyType_name[24:29],
	13:  _KeyType_name[29:34],
	14:  _KeyType_name[34:43],
	15:  _KeyType_name[43:46],
	16:  _KeyType_name[46:47],
	17:  _KeyType_name[47:48],
	18:  _KeyType_name[48:49],
	19:  _KeyType_name[49:50],
	20:  _KeyType_name[50:51],
	21:  _KeyType_name[51:52],
	22:  _KeyType_name[52:53],
	23:  _KeyType_name[53:54],
	24:  _KeyType_name[54:55],
	25:  _KeyType_name[55:56],
	26:  _KeyType_name[56:65],
	27:  _KeyType_name[65:75],
	28:  _KeyType_name[75:80],
	29:  _KeyType_name[80:88],
	30:  _KeyType_name[88:89],
	31:  _KeyType_name[89:90],
	32:  _KeyType_name[90:91],
	33:  _KeyType_name[91:92],
	34:  _KeyType_name[92:93],
	35:  _KeyType_name[93:94],
	36:  _KeyType_name[94:95],
	37:  _KeyType_name[95:96],
	38:  _KeyType_name[96:97],
	39:  _KeyType_name[97:106],
	40:  _KeyType_name[106:116],
	41:  _KeyType_name[116:121],
	42:  _KeyType_name[121:130],
	43:  _KeyType_name[130:139],
	44:  _KeyType_name[139:140],
	45:  _KeyType_name[140:141],
	46:  _KeyType_name[141:142],
	47:  _KeyType_name[142:143],
	48:  _KeyType_name[143:144],
	49:  _KeyType_name[144:145],
	50:  _KeyType_name[145:146],
	51:  _KeyType_name[146:151],
	52:  _KeyType_name[151:154],
	53:  _KeyType_name[154:159],
	54:  _KeyType_name[159:169],
	55:  _KeyType_name[169:179],
	56:  _KeyType_name[179:186],
	57:  _KeyType_name[186:191],
	58:  _KeyType_name[191:199],
	59:  _KeyType_name[199:201],
	60:  _KeyType_name[201:203],
	61:  _KeyType_name[203:205],
	62:  _KeyType_name[205:207],
	63:  _KeyType_name[207:209],
	64:  _KeyType_name[209:211],
	65:  _KeyType_name[211:213],
	66:  _KeyType_name[213:215],
	67:  _KeyType_name[215:217],
	68:  _KeyType_name[217:220],
	69:  _KeyType_name[220:227],
	70:  _KeyType_name[227:237],
	71:  _KeyType_name[237:240],
	72:  _KeyType_name[240:243],
	73:  _KeyType_name[243:246],
	74:  _KeyType_name[246:253],
	75:  _KeyType_name[253:256],
	76:  _KeyType_name[256:259],
	77:  _KeyType_name[259:262],
	78:  _KeyType_name[262:268],
	79:  _KeyType_name[268:271],
	80:  _KeyType_name[271:274],
	81:  _KeyType_name[274:277],
	82:  _KeyType_name[277:280],
	83:  _KeyType_name[280:285],
	85:  _KeyType_name[285:299],
	86:  _KeyType_name[299:304],
	87:  _KeyType_name[304:307],
	88:  _KeyType_name[307:310],
	89:  _KeyType_name[310:312],
	90:  _KeyType_name[312:320],
	91:  _KeyType_name[320:328],
	92:  _KeyType_name[328:334],
	93:  _KeyType_name[334:350],
	94:  _KeyType_name[350:358],
	95:  _KeyType_name[358:367],
	96:  _KeyType_name[367:374],
	97:  _KeyType_name[374:383],
	98:  _KeyType_name[383:390],
	99:  _KeyType_name[390:395],
	100: _KeyType_name[395:403],
	101: _KeyType_name[403:411],
	102: _KeyType_name[411:415],
	103: _KeyType_name[415:417],
	104: _KeyType_name[417:423],
	105: _KeyType_name[423:427],
	106: _KeyType_name[427:432],
	107: _KeyType_name[432:435],
	108: _KeyType_name[435:439],
	109: _KeyType_name[439:447],
	110: _KeyType_name[447:453],
	111: _KeyType_name[453:459],
	112: _KeyType_name[459:464],
	113: _KeyType_name[464:468],
	114: _KeyType_name[468:478],
	115: _KeyType_name[478:486],
	116: _KeyType_name[486:491],
	117: _KeyType_name[491:498],
	118: _KeyType_name[498:509],
	119: _KeyType_name[509:514],
	120: _KeyType_name[514:519],
	121: _KeyType_name[519:526],
	122: _KeyType_name[526:532],
	123: _KeyType_name[532:537],
	124: _KeyType_name[537:540],
	125: _KeyType_name[540:548],
	126: _KeyType_name[548:557],
	127: _KeyType_name[557:564],
	128: _KeyType_name[564:568],
	129: _KeyType_name[568:573],
	130: _KeyType_name[573:578],
	131: _KeyType_name[578:582],
	132: _KeyType_name[582:587],
	133: _KeyType_name[587:591],
	134: _KeyType_name[591:595],
	135: _KeyType_name[595:600],
	136: _KeyType_name[600:604],
	137: _KeyType_name[604:607],
	138: _KeyType_name[607:611],
	139: _KeyType_name[611:615],
	140: _KeyType_name[615:619],
	141: _KeyType_name[619:624],
	142: _KeyType_name[624:629],
	143: _KeyType_name[629:635],
	144: _KeyType_name[635:639],
	145: _KeyType_name[639:647],
	146: _KeyType_name[647:657],
	147: _KeyType_name[657:661],
	148: _KeyType_name[661:666],
	149: _KeyType_name[666:671],
	150: _KeyType_name[671:674],
	151: _KeyType_name[674:679],
	152: _KeyType_name[679:689],
	153: _KeyType_name[689:698],
	154: _KeyType_name[698:710],
	155: _KeyType_name[710:714],
	156: _KeyType_name[714:723],
	157: _KeyType_name[723:731],
	158: _KeyType_name[731:735],
	159: _KeyType_name[735:742],
	160: _KeyType_name[742:749],
	161: _KeyType_name[749:756],
	162: _KeyType_name[756:768],
	163: _KeyType_name[768:776],
	164: _KeyType_name[776:785],
	165: _KeyType_name[785:797],
	166: _KeyType_name[797:803],
	167: _KeyType_name[803:809],
	168: _KeyType_name[809:815],
	169: _KeyType_name[815:820],
	170: _KeyType_name[820:823],
	171: _KeyType_name[823:829],
	172: _KeyType_name[829:837],
	173: _KeyType_name[837:844],
	174: _KeyType_name[844:848],
	175: _KeyType_name[848:852],
	176: _KeyType_name[852:856],
	177: _KeyType_name[856:864],
	178: _KeyType_name[864:874],
	179: _KeyType_name[874:885],
	180: _KeyType_name[885:897],
	181: _KeyType_name[897:900],
	182: _KeyType_name[900:904],
	183: _KeyType_name[904:907],
	184: _KeyType_name[907:910],
	185: _KeyType_name[910:913],
	186: _KeyType_name[913:916],
	187: _KeyType_name[916:919],
	188: _KeyType_name[919:922],
	189: _KeyType_name[922:925],
	190: _KeyType_name[925:928],
	191: _KeyType_name[928:931],
	192: _KeyType_name[931:934],
	193: _KeyType_name[934:937],
	194: _KeyType_name[937:940],
	200: _KeyType_name[940:946],
	201: _KeyType_name[946:953],
	202: _KeyType_name[953:958],
	203: _KeyType_name[958:963],
	204: _KeyType_name[963:972],
	205: _KeyType_name[972:979],
	206: _KeyType_name[979:984],
	207: _KeyType_name[984:988],
	208: _KeyType_name[988:999],
	209: _KeyType_name[999:1008],
	210: _KeyType_name[1008:1013],
	211: _KeyType_name[1013:1015],
	212: _KeyType_name[1015:1021],
	213: _KeyType_name[1021:1026],
	214: _KeyType_name[1026:1034],
	215: _KeyType_name[1034:1039],
	216: _KeyType_name[1039:1043],
	217: _KeyType_name[1043:1049],
	218: _KeyType_name[1049:1056],
	219: _KeyType_name[1056:1063],
	220: _KeyType_name[1063:1068],
	221: _KeyType_name[1068:1072],
	222: _KeyType_name[1072:1080],
	223: _KeyType_name[1080:1086],
	224: _KeyType_name[1086:1100],
	225: _KeyType_name[1100:1112],
	226: _KeyType_name[1112:1117],
	227: _KeyType_name[1117:1132],
	228: _KeyType_name[1132:1146],
	229: _KeyType_name[1146:1158],
	230: _KeyType_name[1158:1168],
	231: _KeyType_name[1168:1172],
	232: _KeyType_name[1172:1177],
	233: _KeyType_name[1177:1188],
	234: _KeyType_name[1188:1192],
	235: _KeyType_name[1192:1201],
	236: _KeyType_name[1201:1208],
	237: _KeyType_name[1208:1217],
	238: _KeyType_name[1217:1221],
	239: _KeyType_name[1221:1224],
	240: _KeyType_name[1224:1231],
	241: _KeyType_name[1231:1240],
	242: _KeyType_name[1240:1253],
	243: _KeyType_name[1253:1268],
	244: _KeyType_name[1268:1282],
	245: _KeyType_name[1282:1292],
	246: _KeyType_name[1292:1297],
	247: _KeyType_name[1297:1303],
	248: _KeyType_name[1303:1310],
	256: _KeyType_name[1310:1314],
	257: _KeyType_name[1314:1318],
	258: _KeyType_name[1318:1322],
	259: _KeyType_name[1322:1326],
	260: _KeyType_name[1326:1330],
	261: _KeyType_name[1330:1334],
	262: _KeyType_name[1334:1338],
	263: _KeyType_name[1338:1342],
	264: _KeyType_name[1342:1346],
	265: _KeyType_name[1346:1350],
	272: _KeyType_name[1350:1358],
	273: _KeyType_name[1358:1366],
	274: _KeyType_name[1366:1375],
	275: _KeyType_name[1375:1382],
	276: _KeyType_name[1382:1390],
	277: _KeyType_name[1390:1400],
	278: _KeyType_name[1400:1407],
	279: _KeyType_name[1407:1414],
	288: _KeyType_name[1414:1425],
	289: _KeyType_name[1425:1433],
	290: _KeyType_name[1433:1442],
	291: _KeyType_name[1442:1448],
	292: _KeyType_name[1448:1455],
	293: _KeyType_name[1455:1464],
	294: _KeyType_name[1464:1471],
	295: _KeyType_name[1471:1479],
	296: _KeyType_name[1479:1487],
	297: _KeyType_name[1487:1495],
	298: _KeyType_name[1495:1503],
	299: _KeyType_name[1503:1511],
	303: _KeyType_name[1511:1518],
	304: _KeyType_name[1518:1528],
	305: _KeyType_name[1528:1532],
	306: _KeyType_name[1532:1536],
	307: _KeyType_name[1536:1540],
	308: _KeyType_name[1540:1544],
	309: _KeyType_name[1544:1548],
	310: _KeyType_name[1548:1553],
	311: _KeyType_name[1553:1558],
	312: _KeyType_name[1558:1564],
	313: _KeyType_name[1564:1570],
	314: _KeyType_name[1570:1579],
	315: _KeyType_name[1579:1587],
	316: _KeyType_name[1587:1594],
	317: _KeyType_name[1594:1603],
	318: _KeyType_name[1603:1612],
	320: _KeyType_name[1612:1619],
	321: _KeyType_name[1619:1632],
	322: _KeyType_name[1632:1644],
	323: _KeyType_name[1644:1657],
	324: _KeyType_name[1657:1672],
	325: _KeyType_name[1672:1685],
	326: _KeyType_name[1685:1697],
	327: _KeyType_name[1697:1708],
	328: _KeyType_name[1708:1723],
	330: _KeyType_name[1723:1731],
	331: _KeyType_name[1731:1740],
	332: _KeyType_name[1740:1750],
	333: _KeyType_name[1750:1766],
	334: _KeyType_name[1766:1783],
	335: _KeyType_name[1783:1797],
	336: _KeyType_name[1797:1805],
	337: _KeyType_name[1805:1814],
	352: _KeyType_name[1814:1816],
	353: _KeyType_name[1816:1822],
	354: _KeyType_name[1822:1826],
	355: _KeyType_name[1826:1831],
	356: _KeyType_name[1831:1837],
	357: _KeyType_name[1837:1843],
	358: _KeyType_name[1843:1847],
	359: _KeyType_name[1847:1851],
	360: _KeyType_name[1851:1857],
	361: _KeyType_name[1857:1864],
	362: _KeyType_name[1864:1871],
	363: _KeyType_name[1871:1878],
	364: _KeyType_name[1878:1887],
	365: _KeyType_name[1887:1890],
	366: _KeyType_name[1890:1893],
	367: _KeyType_name[1893:1896],
	368: _KeyType_name[1896:1904],
	369: _KeyType_name[1904:1909],
	370: _KeyType_name[1909:1917],
	371: _KeyType_name[1917:1922],
	372: _KeyType_name[1922:1926],
	373: _KeyType_name[1926:1930],
	374: _KeyType_name[1930:1938],
	375: _KeyType_name[1938:1944],
	376: _KeyType_name[1944:1946],
	377: _KeyType_name[1946:1948],
	378: _KeyType_name[1948:1951],
	379: _KeyType_name[1951:1954],
	380: _KeyType_name[1954:1958],
	381: _KeyType_name[1958:1961],
	382: _KeyType_name[1961:1965],
	383: _KeyType_name[1965:1967],
	384: _KeyType_name[1967:1971],
	385: _KeyType_name[1971:1976],
	386: _KeyType_name[1976:1981],
	387: _KeyType_name[1981:1987],
	388: _KeyType_name[1987:1991],
	389: _KeyType_name[1991:1994],
	390: _KeyType_name[1994:1997],
	391: _KeyType_name[1997:2000],
	392: _KeyType_name[2000:2005],
	393: _KeyType_name[2005:2010],
	394: _KeyType_name[2010:2019],
	395: _KeyType_name[2019:2023],
	396: _KeyType_name[2023:2027],
	397: _KeyType_name[2027:2035],
	398: _KeyType_name[2035:2038],
	399: _KeyType_name[2038:2043],
	400: _KeyType_name[2043:2049],
	401: _KeyType_name[2049:2053],
	402: _KeyType_name[2053:2062],
	403: _KeyType_name[2062:2073],
	404: _KeyType_name[2073:2078],
	405: _KeyType_name[2078:2082],
	406: _KeyType_name[2082:2084],
	407: _KeyType_name[2084:2088],
	408: _KeyType_name[2088:2095],
	409: _KeyType_name[2095:2099],
	410: _KeyType_name[2099:2106],
	411: _KeyType_name[2106:2111],
	412: _KeyType_name[2111:2119],
	413: _KeyType_name[2119:2125],
	414: _KeyType_name[2125:2129],
	415: _KeyType_name[2129:2133],
	416: _KeyType_name[2133:2143],
	417: _KeyType_name[2143:2148],
	418: _KeyType_name[2148:2154],
	419: _KeyType_name[2154:2161],
	420: _KeyType_name[2161:2170],
	421: _KeyType_name[2170:2183],
	422: _KeyType_name[2183:2189],
	423: _KeyType_name[2189:2200],
	424: _KeyType_name[2200:2214],
	425: _KeyType_name[2214:2226],
	426: _KeyType_name[2226:2234],
	427: _KeyType_name[2234:2238],
	428: _KeyType_name[2238:2247],
	429: _KeyType_name[2247:2258],
	430: _KeyType_name[2258:2267],
	431: _KeyType_name[2267:2280],
	432: _KeyType_name[2280:2290],
	433: _KeyType_name[2290:2296],
	434: _KeyType_name[2296:2302],
	435: _KeyType_name[2302:2306],
	436: _KeyType_name[2306:2315],
	437: _KeyType_name[2315:2327],
	438: _KeyType_name[2327:2338],
	439: _KeyType_name[2338:2349],
	440: _KeyType_name[2349:2361],
	441: _KeyType_name[2361:2375],
	442: _KeyType_name[2375:2381],
	448: _KeyType_name[2381:2387],
	449: _KeyType_name[2387:2393],
	450: _KeyType_name[2393:2400],
	451: _KeyType_name[2400:2407],
	464: _KeyType_name[2407:2409],
	465: _KeyType_name[2409:2414],
	466: _KeyType_name[2414:2418],
	467: _KeyType_name[2418:2422],
	468: _KeyType_name[2422:2426],
	469: _KeyType_name[2426:2430],
	470: _KeyType_name[2430:2434],
	471: _KeyType_name[2434:2438],
	472: _KeyType_name[2438:2442],
	473: _KeyType_name[2442:2446],
	474: _KeyType_name[2446:2450],
	475: _KeyType_name[2450:2455],
	476: _KeyType_name[2455:2460],
	477: _KeyType_name[2460:2465],
	478: _KeyType_name[2465:2468],
	479: _KeyType_name[2468:2471],
	480: _KeyType_name[2471:2474],
	481: _KeyType_name[2474:2477],
	482: _KeyType_name[2477:2480],
	483: _KeyType_name[2480:2483],
	484: _KeyType_name[2483:2486],
	497: _KeyType_name[2486:2493],
	498: _KeyType_name[2493:2500],
	499: _KeyType_name[2500:2507],
	500: _KeyType_name[2507:2514],
	501: _KeyType_name[2514:2521],
	502: _KeyType_name[2521:2528],
	503: _KeyType_name[2528:2535],
	504: _KeyType_name[2535:2542],
	505: _KeyType_name[2542:2549],
	506: _KeyType_name[2549:2557],
	512: _KeyType_name[2557:2565],
	513: _KeyType_name[2565:2573],
	514: _KeyType_name[2573:2581],
	515: _KeyType_name[2581:2589],
	516: _KeyType_name[2589:2597],
	517: _KeyType_name[2597:2605],
	518: _KeyType_name[2605:2613],
	519: _KeyType_name[2613:2621],
	520: _KeyType_name[2621:2629],
	521: _KeyType_name[2629:2637],
	522: _KeyType_name[2637:2648],
	523: _KeyType_name[2648:2660],
	528: _KeyType_name[2660:2671],
	529: _KeyType_name[2671:2680],
	530: _KeyType_name[2680:2694],
	531: _KeyType_name[2694:2704],
	532: _KeyType_name[2704:2715],
	533: _KeyType_name[2715:2727],
	534: _KeyType_name[2727:2740],
	535: _KeyType_name[2740:2748],
	536: _KeyType_name[2748:2758],
	537: _KeyType_name[2758:2768],
	538: _KeyType_name[2768:2779],
	539: _KeyType_name[2779:2790],
	540: _KeyType_name[2790:2802],
	541: _KeyType_name[2802:2817],
	542: _KeyType_name[2817:2829],
	704: _KeyType_name[2829:2844],
	705: _KeyType_name[2844:2860],
	706: _KeyType_name[2860:2876],
	707: _KeyType_name[2876:2892],
	708: _KeyType_name[2892:2908],
	709: _KeyType_name[2908:2924],
	710: _KeyType_name[2924:2940],
	711: _KeyType_name[2940:2956],
	712: _KeyType_name[2956:2972],
	713: _KeyType_name[2972:2989],
	714: _KeyType_name[2989:3006],
	715: _KeyType_name[3006:3023],
	716: _KeyType_name[3023:3040],
	717: _KeyType_name[3040:3057],
	718: _KeyType_name[3057:3074],
	719: _KeyType_name[3074:3091],
	720: _KeyType_name[3091:3108],
	721: _KeyType_name[3108:3125],
	722: _KeyType_name[3125:3142],
	723: _KeyType_name[3142:3159],
	724: _KeyType_name[3159:3176],
	725: _KeyType_name[3176:3193],
	726: _KeyType_name[3193:3210],
	727: _KeyType_name[3210:3227],
	728: _KeyType_name[3227:3244],
	729: _KeyType_name[3244:3261],
	730: _KeyType_name[3261:3278],
	731: _KeyType_name[3278:3295],
	732: _KeyType_name[3295:3312],
	733: _KeyType_name[3312:3329],
	734: _KeyType_name[3329:3346],
	735: _KeyType_name[3346:3363],
	736: _KeyType_name[3363:3380],
	737: _KeyType_name[3380:3397],
	738: _KeyType_name[3397:3414],
	739: _KeyType_name[3414:3431],
	740: _KeyType_name[3431:3448],
	741: _KeyType_name[3448:3465],
	742: _KeyType_name[3465:3482],
	743: _KeyType_name[3482:3499],
}

func (i KeyType) String() string {
	if str, ok := _KeyType_map[i]; ok {
		return str
	}
	return "KeyType(" + strconv.FormatInt(int64(i), 10) + ")"
}

const _LEDType_name = "NumLockCapsLockScrollLockComposeKanaSleepSuspendMuteMiscMailCharging"

var _LEDType_index = [...]uint8{0, 7, 15, 25, 32, 36, 41, 48, 52, 56, 60, 68}

func (i LEDType) String() string {
	if i < 0 || i >= LEDType(len(_LEDType_index)-1) {
		return "LEDType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LEDType_name[_LEDType_index[i]:_LEDType_index[i+1]]
}

const _MiscType_name = "SerialPulseLedGestureRawScanTimestamp"

var _MiscType_index = [...]uint8{0, 6, 14, 21, 24, 28, 37}

func (i MiscType) String() string {
	if i < 0 || i >= MiscType(len(_MiscType_index)-1) {
		return "MiscType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MiscType_name[_MiscType_index[i]:_MiscType_index[i+1]]
}

const _MtToolType_name = "FingerPen"

var _MtToolType_index = [...]uint8{0, 6, 9}

func (i MtToolType) String() string {
	if i < 0 || i >= MtToolType(len(_MtToolType_index)-1) {
		return "MtToolType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MtToolType_name[_MtToolType_index[i]:_MtToolType_index[i+1]]
}

const _PowerType_name = "OffOnStandby"

var _PowerType_index = [...]uint8{0, 3, 5, 12}

func (i PowerType) String() string {
	if i < 0 || i >= PowerType(len(_PowerType_index)-1) {
		return "PowerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PowerType_name[_PowerType_index[i]:_PowerType_index[i+1]]
}

const _RelativeType_name = "XYZRXRYRZHWheelDialWheelMisc"

var _RelativeType_index = [...]uint8{0, 1, 2, 3, 5, 7, 9, 15, 19, 24, 28}

func (i RelativeType) String() string {
	if i < 0 || i >= RelativeType(len(_RelativeType_index)-1) {
		return "RelativeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RelativeType_name[_RelativeType_index[i]:_RelativeType_index[i+1]]
}

const _RepeatType_name = "DelayPeriod"

var _RepeatType_index = [...]uint8{0, 5, 11}

func (i RepeatType) String() string {
	if i < 0 || i >= RepeatType(len(_RepeatType_index)-1) {
		return "RepeatType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RepeatType_name[_RepeatType_index[i]:_RepeatType_index[i+1]]
}

const _SoundType_name = "ClickBellTone"

var _SoundType_index = [...]uint8{0, 5, 9, 13}

func (i SoundType) String() string {
	if i < 0 || i >= SoundType(len(_SoundType_index)-1) {
		return "SoundType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SoundType_name[_SoundType_index[i]:_SoundType_index[i+1]]
}

const _SwitchType_name = "LidTabletModeHeadphoneInsertRFKillAllMicrophoneInsertDockLineoutInsertJackPhysicalInsertVideoOutInsertCameraLensCoverKeypadSlideFrontProximityRotateLockLineInInsert"

var _SwitchType_index = [...]uint8{0, 3, 13, 28, 37, 53, 57, 70, 88, 102, 117, 128, 142, 152, 164}

func (i SwitchType) String() string {
	if i < 0 || i >= SwitchType(len(_SwitchType_index)-1) {
		return "SwitchType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SwitchType_name[_SwitchType_index[i]:_SwitchType_index[i+1]]
}

const _SyncType_name = "ReportConfigMTReportDropped"

var _SyncType_index = [...]uint8{0, 6, 12, 20, 27}

func (i SyncType) String() string {
	if i < 0 || i >= SyncType(len(_SyncType_index)-1) {
		return "SyncType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SyncType_name[_SyncType_index[i]:_SyncType_index[i+1]]
}
