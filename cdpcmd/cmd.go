// Code generated by cdpgen; DO NOT EDIT!

package cdpcmd

// CmdType is the type for CDP methods names.
type CmdType string

func (c CmdType) String() string {
	return string(c)
}

// Cmd methods.
const (
	AccessibilityGetPartialAXTree                   CmdType = "Accessibility.getPartialAXTree"
	AnimationEnable                                 CmdType = "Animation.enable"
	AnimationDisable                                CmdType = "Animation.disable"
	AnimationGetPlaybackRate                        CmdType = "Animation.getPlaybackRate"
	AnimationSetPlaybackRate                        CmdType = "Animation.setPlaybackRate"
	AnimationGetCurrentTime                         CmdType = "Animation.getCurrentTime"
	AnimationSetPaused                              CmdType = "Animation.setPaused"
	AnimationSetTiming                              CmdType = "Animation.setTiming"
	AnimationSeekAnimations                         CmdType = "Animation.seekAnimations"
	AnimationReleaseAnimations                      CmdType = "Animation.releaseAnimations"
	AnimationResolveAnimation                       CmdType = "Animation.resolveAnimation"
	ApplicationCacheGetFramesWithManifests          CmdType = "ApplicationCache.getFramesWithManifests"
	ApplicationCacheEnable                          CmdType = "ApplicationCache.enable"
	ApplicationCacheGetManifestForFrame             CmdType = "ApplicationCache.getManifestForFrame"
	ApplicationCacheGetApplicationCacheForFrame     CmdType = "ApplicationCache.getApplicationCacheForFrame"
	BrowserGetWindowForTarget                       CmdType = "Browser.getWindowForTarget"
	BrowserSetWindowBounds                          CmdType = "Browser.setWindowBounds"
	BrowserGetWindowBounds                          CmdType = "Browser.getWindowBounds"
	CSSEnable                                       CmdType = "CSS.enable"
	CSSDisable                                      CmdType = "CSS.disable"
	CSSGetMatchedStylesForNode                      CmdType = "CSS.getMatchedStylesForNode"
	CSSGetInlineStylesForNode                       CmdType = "CSS.getInlineStylesForNode"
	CSSGetComputedStyleForNode                      CmdType = "CSS.getComputedStyleForNode"
	CSSGetPlatformFontsForNode                      CmdType = "CSS.getPlatformFontsForNode"
	CSSGetStyleSheetText                            CmdType = "CSS.getStyleSheetText"
	CSSCollectClassNames                            CmdType = "CSS.collectClassNames"
	CSSSetStyleSheetText                            CmdType = "CSS.setStyleSheetText"
	CSSSetRuleSelector                              CmdType = "CSS.setRuleSelector"
	CSSSetKeyframeKey                               CmdType = "CSS.setKeyframeKey"
	CSSSetStyleTexts                                CmdType = "CSS.setStyleTexts"
	CSSSetMediaText                                 CmdType = "CSS.setMediaText"
	CSSCreateStyleSheet                             CmdType = "CSS.createStyleSheet"
	CSSAddRule                                      CmdType = "CSS.addRule"
	CSSForcePseudoState                             CmdType = "CSS.forcePseudoState"
	CSSGetMediaQueries                              CmdType = "CSS.getMediaQueries"
	CSSSetEffectivePropertyValueForNode             CmdType = "CSS.setEffectivePropertyValueForNode"
	CSSGetBackgroundColors                          CmdType = "CSS.getBackgroundColors"
	CSSGetLayoutTreeAndStyles                       CmdType = "CSS.getLayoutTreeAndStyles"
	CSSStartRuleUsageTracking                       CmdType = "CSS.startRuleUsageTracking"
	CSSTakeCoverageDelta                            CmdType = "CSS.takeCoverageDelta"
	CSSStopRuleUsageTracking                        CmdType = "CSS.stopRuleUsageTracking"
	CacheStorageRequestCacheNames                   CmdType = "CacheStorage.requestCacheNames"
	CacheStorageRequestEntries                      CmdType = "CacheStorage.requestEntries"
	CacheStorageDeleteCache                         CmdType = "CacheStorage.deleteCache"
	CacheStorageDeleteEntry                         CmdType = "CacheStorage.deleteEntry"
	ConsoleEnable                                   CmdType = "Console.enable"
	ConsoleDisable                                  CmdType = "Console.disable"
	ConsoleClearMessages                            CmdType = "Console.clearMessages"
	DOMEnable                                       CmdType = "DOM.enable"
	DOMDisable                                      CmdType = "DOM.disable"
	DOMGetDocument                                  CmdType = "DOM.getDocument"
	DOMGetFlattenedDocument                         CmdType = "DOM.getFlattenedDocument"
	DOMCollectClassNamesFromSubtree                 CmdType = "DOM.collectClassNamesFromSubtree"
	DOMRequestChildNodes                            CmdType = "DOM.requestChildNodes"
	DOMQuerySelector                                CmdType = "DOM.querySelector"
	DOMQuerySelectorAll                             CmdType = "DOM.querySelectorAll"
	DOMSetNodeName                                  CmdType = "DOM.setNodeName"
	DOMSetNodeValue                                 CmdType = "DOM.setNodeValue"
	DOMRemoveNode                                   CmdType = "DOM.removeNode"
	DOMSetAttributeValue                            CmdType = "DOM.setAttributeValue"
	DOMSetAttributesAsText                          CmdType = "DOM.setAttributesAsText"
	DOMRemoveAttribute                              CmdType = "DOM.removeAttribute"
	DOMGetOuterHTML                                 CmdType = "DOM.getOuterHTML"
	DOMSetOuterHTML                                 CmdType = "DOM.setOuterHTML"
	DOMPerformSearch                                CmdType = "DOM.performSearch"
	DOMGetSearchResults                             CmdType = "DOM.getSearchResults"
	DOMDiscardSearchResults                         CmdType = "DOM.discardSearchResults"
	DOMRequestNode                                  CmdType = "DOM.requestNode"
	DOMSetInspectMode                               CmdType = "DOM.setInspectMode"
	DOMHighlightRect                                CmdType = "DOM.highlightRect"
	DOMHighlightQuad                                CmdType = "DOM.highlightQuad"
	DOMHighlightNode                                CmdType = "DOM.highlightNode"
	DOMHideHighlight                                CmdType = "DOM.hideHighlight"
	DOMHighlightFrame                               CmdType = "DOM.highlightFrame"
	DOMPushNodeByPathToFrontend                     CmdType = "DOM.pushNodeByPathToFrontend"
	DOMPushNodesByBackendIdsToFrontend              CmdType = "DOM.pushNodesByBackendIdsToFrontend"
	DOMSetInspectedNode                             CmdType = "DOM.setInspectedNode"
	DOMResolveNode                                  CmdType = "DOM.resolveNode"
	DOMGetAttributes                                CmdType = "DOM.getAttributes"
	DOMCopyTo                                       CmdType = "DOM.copyTo"
	DOMMoveTo                                       CmdType = "DOM.moveTo"
	DOMUndo                                         CmdType = "DOM.undo"
	DOMRedo                                         CmdType = "DOM.redo"
	DOMMarkUndoableState                            CmdType = "DOM.markUndoableState"
	DOMFocus                                        CmdType = "DOM.focus"
	DOMSetFileInputFiles                            CmdType = "DOM.setFileInputFiles"
	DOMGetBoxModel                                  CmdType = "DOM.getBoxModel"
	DOMGetNodeForLocation                           CmdType = "DOM.getNodeForLocation"
	DOMGetRelayoutBoundary                          CmdType = "DOM.getRelayoutBoundary"
	DOMGetHighlightObjectForTest                    CmdType = "DOM.getHighlightObjectForTest"
	DOMDebuggerSetDOMBreakpoint                     CmdType = "DOMDebugger.setDOMBreakpoint"
	DOMDebuggerRemoveDOMBreakpoint                  CmdType = "DOMDebugger.removeDOMBreakpoint"
	DOMDebuggerSetEventListenerBreakpoint           CmdType = "DOMDebugger.setEventListenerBreakpoint"
	DOMDebuggerRemoveEventListenerBreakpoint        CmdType = "DOMDebugger.removeEventListenerBreakpoint"
	DOMDebuggerSetInstrumentationBreakpoint         CmdType = "DOMDebugger.setInstrumentationBreakpoint"
	DOMDebuggerRemoveInstrumentationBreakpoint      CmdType = "DOMDebugger.removeInstrumentationBreakpoint"
	DOMDebuggerSetXHRBreakpoint                     CmdType = "DOMDebugger.setXHRBreakpoint"
	DOMDebuggerRemoveXHRBreakpoint                  CmdType = "DOMDebugger.removeXHRBreakpoint"
	DOMDebuggerGetEventListeners                    CmdType = "DOMDebugger.getEventListeners"
	DOMStorageEnable                                CmdType = "DOMStorage.enable"
	DOMStorageDisable                               CmdType = "DOMStorage.disable"
	DOMStorageClear                                 CmdType = "DOMStorage.clear"
	DOMStorageGetDOMStorageItems                    CmdType = "DOMStorage.getDOMStorageItems"
	DOMStorageSetDOMStorageItem                     CmdType = "DOMStorage.setDOMStorageItem"
	DOMStorageRemoveDOMStorageItem                  CmdType = "DOMStorage.removeDOMStorageItem"
	DatabaseEnable                                  CmdType = "Database.enable"
	DatabaseDisable                                 CmdType = "Database.disable"
	DatabaseGetDatabaseTableNames                   CmdType = "Database.getDatabaseTableNames"
	DatabaseExecuteSQL                              CmdType = "Database.executeSQL"
	DebuggerEnable                                  CmdType = "Debugger.enable"
	DebuggerDisable                                 CmdType = "Debugger.disable"
	DebuggerSetBreakpointsActive                    CmdType = "Debugger.setBreakpointsActive"
	DebuggerSetSkipAllPauses                        CmdType = "Debugger.setSkipAllPauses"
	DebuggerSetBreakpointByURL                      CmdType = "Debugger.setBreakpointByUrl"
	DebuggerSetBreakpoint                           CmdType = "Debugger.setBreakpoint"
	DebuggerRemoveBreakpoint                        CmdType = "Debugger.removeBreakpoint"
	DebuggerGetPossibleBreakpoints                  CmdType = "Debugger.getPossibleBreakpoints"
	DebuggerContinueToLocation                      CmdType = "Debugger.continueToLocation"
	DebuggerStepOver                                CmdType = "Debugger.stepOver"
	DebuggerStepInto                                CmdType = "Debugger.stepInto"
	DebuggerStepOut                                 CmdType = "Debugger.stepOut"
	DebuggerPause                                   CmdType = "Debugger.pause"
	DebuggerScheduleStepIntoAsync                   CmdType = "Debugger.scheduleStepIntoAsync"
	DebuggerResume                                  CmdType = "Debugger.resume"
	DebuggerSearchInContent                         CmdType = "Debugger.searchInContent"
	DebuggerSetScriptSource                         CmdType = "Debugger.setScriptSource"
	DebuggerRestartFrame                            CmdType = "Debugger.restartFrame"
	DebuggerGetScriptSource                         CmdType = "Debugger.getScriptSource"
	DebuggerSetPauseOnExceptions                    CmdType = "Debugger.setPauseOnExceptions"
	DebuggerEvaluateOnCallFrame                     CmdType = "Debugger.evaluateOnCallFrame"
	DebuggerSetVariableValue                        CmdType = "Debugger.setVariableValue"
	DebuggerSetAsyncCallStackDepth                  CmdType = "Debugger.setAsyncCallStackDepth"
	DebuggerSetBlackboxPatterns                     CmdType = "Debugger.setBlackboxPatterns"
	DebuggerSetBlackboxedRanges                     CmdType = "Debugger.setBlackboxedRanges"
	DeviceOrientationSetDeviceOrientationOverride   CmdType = "DeviceOrientation.setDeviceOrientationOverride"
	DeviceOrientationClearDeviceOrientationOverride CmdType = "DeviceOrientation.clearDeviceOrientationOverride"
	EmulationSetDeviceMetricsOverride               CmdType = "Emulation.setDeviceMetricsOverride"
	EmulationClearDeviceMetricsOverride             CmdType = "Emulation.clearDeviceMetricsOverride"
	EmulationForceViewport                          CmdType = "Emulation.forceViewport"
	EmulationResetViewport                          CmdType = "Emulation.resetViewport"
	EmulationResetPageScaleFactor                   CmdType = "Emulation.resetPageScaleFactor"
	EmulationSetPageScaleFactor                     CmdType = "Emulation.setPageScaleFactor"
	EmulationSetVisibleSize                         CmdType = "Emulation.setVisibleSize"
	EmulationSetScriptExecutionDisabled             CmdType = "Emulation.setScriptExecutionDisabled"
	EmulationSetGeolocationOverride                 CmdType = "Emulation.setGeolocationOverride"
	EmulationClearGeolocationOverride               CmdType = "Emulation.clearGeolocationOverride"
	EmulationSetTouchEmulationEnabled               CmdType = "Emulation.setTouchEmulationEnabled"
	EmulationSetEmulatedMedia                       CmdType = "Emulation.setEmulatedMedia"
	EmulationSetCPUThrottlingRate                   CmdType = "Emulation.setCPUThrottlingRate"
	EmulationCanEmulate                             CmdType = "Emulation.canEmulate"
	EmulationSetVirtualTimePolicy                   CmdType = "Emulation.setVirtualTimePolicy"
	EmulationSetDefaultBackgroundColorOverride      CmdType = "Emulation.setDefaultBackgroundColorOverride"
	HeapProfilerEnable                              CmdType = "HeapProfiler.enable"
	HeapProfilerDisable                             CmdType = "HeapProfiler.disable"
	HeapProfilerStartTrackingHeapObjects            CmdType = "HeapProfiler.startTrackingHeapObjects"
	HeapProfilerStopTrackingHeapObjects             CmdType = "HeapProfiler.stopTrackingHeapObjects"
	HeapProfilerTakeHeapSnapshot                    CmdType = "HeapProfiler.takeHeapSnapshot"
	HeapProfilerCollectGarbage                      CmdType = "HeapProfiler.collectGarbage"
	HeapProfilerGetObjectByHeapObjectID             CmdType = "HeapProfiler.getObjectByHeapObjectId"
	HeapProfilerAddInspectedHeapObject              CmdType = "HeapProfiler.addInspectedHeapObject"
	HeapProfilerGetHeapObjectID                     CmdType = "HeapProfiler.getHeapObjectId"
	HeapProfilerStartSampling                       CmdType = "HeapProfiler.startSampling"
	HeapProfilerStopSampling                        CmdType = "HeapProfiler.stopSampling"
	IORead                                          CmdType = "IO.read"
	IOClose                                         CmdType = "IO.close"
	IndexedDBEnable                                 CmdType = "IndexedDB.enable"
	IndexedDBDisable                                CmdType = "IndexedDB.disable"
	IndexedDBRequestDatabaseNames                   CmdType = "IndexedDB.requestDatabaseNames"
	IndexedDBRequestDatabase                        CmdType = "IndexedDB.requestDatabase"
	IndexedDBRequestData                            CmdType = "IndexedDB.requestData"
	IndexedDBClearObjectStore                       CmdType = "IndexedDB.clearObjectStore"
	IndexedDBDeleteDatabase                         CmdType = "IndexedDB.deleteDatabase"
	InputDispatchKeyEvent                           CmdType = "Input.dispatchKeyEvent"
	InputDispatchMouseEvent                         CmdType = "Input.dispatchMouseEvent"
	InputDispatchTouchEvent                         CmdType = "Input.dispatchTouchEvent"
	InputEmulateTouchFromMouseEvent                 CmdType = "Input.emulateTouchFromMouseEvent"
	InputSynthesizePinchGesture                     CmdType = "Input.synthesizePinchGesture"
	InputSynthesizeScrollGesture                    CmdType = "Input.synthesizeScrollGesture"
	InputSynthesizeTapGesture                       CmdType = "Input.synthesizeTapGesture"
	InspectorEnable                                 CmdType = "Inspector.enable"
	InspectorDisable                                CmdType = "Inspector.disable"
	LayerTreeEnable                                 CmdType = "LayerTree.enable"
	LayerTreeDisable                                CmdType = "LayerTree.disable"
	LayerTreeCompositingReasons                     CmdType = "LayerTree.compositingReasons"
	LayerTreeMakeSnapshot                           CmdType = "LayerTree.makeSnapshot"
	LayerTreeLoadSnapshot                           CmdType = "LayerTree.loadSnapshot"
	LayerTreeReleaseSnapshot                        CmdType = "LayerTree.releaseSnapshot"
	LayerTreeProfileSnapshot                        CmdType = "LayerTree.profileSnapshot"
	LayerTreeReplaySnapshot                         CmdType = "LayerTree.replaySnapshot"
	LayerTreeSnapshotCommandLog                     CmdType = "LayerTree.snapshotCommandLog"
	LogEnable                                       CmdType = "Log.enable"
	LogDisable                                      CmdType = "Log.disable"
	LogClear                                        CmdType = "Log.clear"
	LogStartViolationsReport                        CmdType = "Log.startViolationsReport"
	LogStopViolationsReport                         CmdType = "Log.stopViolationsReport"
	MemoryGetDOMCounters                            CmdType = "Memory.getDOMCounters"
	MemorySetPressureNotificationsSuppressed        CmdType = "Memory.setPressureNotificationsSuppressed"
	MemorySimulatePressureNotification              CmdType = "Memory.simulatePressureNotification"
	NetworkEnable                                   CmdType = "Network.enable"
	NetworkDisable                                  CmdType = "Network.disable"
	NetworkSetUserAgentOverride                     CmdType = "Network.setUserAgentOverride"
	NetworkSetExtraHTTPHeaders                      CmdType = "Network.setExtraHTTPHeaders"
	NetworkGetResponseBody                          CmdType = "Network.getResponseBody"
	NetworkSetBlockedURLs                           CmdType = "Network.setBlockedURLs"
	NetworkReplayXHR                                CmdType = "Network.replayXHR"
	NetworkSetMonitoringXHREnabled                  CmdType = "Network.setMonitoringXHREnabled"
	NetworkCanClearBrowserCache                     CmdType = "Network.canClearBrowserCache"
	NetworkClearBrowserCache                        CmdType = "Network.clearBrowserCache"
	NetworkCanClearBrowserCookies                   CmdType = "Network.canClearBrowserCookies"
	NetworkClearBrowserCookies                      CmdType = "Network.clearBrowserCookies"
	NetworkGetCookies                               CmdType = "Network.getCookies"
	NetworkGetAllCookies                            CmdType = "Network.getAllCookies"
	NetworkDeleteCookie                             CmdType = "Network.deleteCookie"
	NetworkSetCookie                                CmdType = "Network.setCookie"
	NetworkCanEmulateNetworkConditions              CmdType = "Network.canEmulateNetworkConditions"
	NetworkEmulateNetworkConditions                 CmdType = "Network.emulateNetworkConditions"
	NetworkSetCacheDisabled                         CmdType = "Network.setCacheDisabled"
	NetworkSetBypassServiceWorker                   CmdType = "Network.setBypassServiceWorker"
	NetworkSetDataSizeLimitsForTest                 CmdType = "Network.setDataSizeLimitsForTest"
	NetworkGetCertificate                           CmdType = "Network.getCertificate"
	PageEnable                                      CmdType = "Page.enable"
	PageDisable                                     CmdType = "Page.disable"
	PageAddScriptToEvaluateOnLoad                   CmdType = "Page.addScriptToEvaluateOnLoad"
	PageRemoveScriptToEvaluateOnLoad                CmdType = "Page.removeScriptToEvaluateOnLoad"
	PageSetAutoAttachToCreatedPages                 CmdType = "Page.setAutoAttachToCreatedPages"
	PageReload                                      CmdType = "Page.reload"
	PageNavigate                                    CmdType = "Page.navigate"
	PageStopLoading                                 CmdType = "Page.stopLoading"
	PageGetNavigationHistory                        CmdType = "Page.getNavigationHistory"
	PageNavigateToHistoryEntry                      CmdType = "Page.navigateToHistoryEntry"
	PageGetCookies                                  CmdType = "Page.getCookies"
	PageDeleteCookie                                CmdType = "Page.deleteCookie"
	PageGetResourceTree                             CmdType = "Page.getResourceTree"
	PageGetResourceContent                          CmdType = "Page.getResourceContent"
	PageSearchInResource                            CmdType = "Page.searchInResource"
	PageSetDocumentContent                          CmdType = "Page.setDocumentContent"
	PageSetDeviceMetricsOverride                    CmdType = "Page.setDeviceMetricsOverride"
	PageClearDeviceMetricsOverride                  CmdType = "Page.clearDeviceMetricsOverride"
	PageSetGeolocationOverride                      CmdType = "Page.setGeolocationOverride"
	PageClearGeolocationOverride                    CmdType = "Page.clearGeolocationOverride"
	PageSetDeviceOrientationOverride                CmdType = "Page.setDeviceOrientationOverride"
	PageClearDeviceOrientationOverride              CmdType = "Page.clearDeviceOrientationOverride"
	PageSetTouchEmulationEnabled                    CmdType = "Page.setTouchEmulationEnabled"
	PageCaptureScreenshot                           CmdType = "Page.captureScreenshot"
	PagePrintToPDF                                  CmdType = "Page.printToPDF"
	PageStartScreencast                             CmdType = "Page.startScreencast"
	PageStopScreencast                              CmdType = "Page.stopScreencast"
	PageScreencastFrameAck                          CmdType = "Page.screencastFrameAck"
	PageHandleJavaScriptDialog                      CmdType = "Page.handleJavaScriptDialog"
	PageSetColorPickerEnabled                       CmdType = "Page.setColorPickerEnabled"
	PageConfigureOverlay                            CmdType = "Page.configureOverlay"
	PageGetAppManifest                              CmdType = "Page.getAppManifest"
	PageRequestAppBanner                            CmdType = "Page.requestAppBanner"
	PageSetControlNavigations                       CmdType = "Page.setControlNavigations"
	PageProcessNavigation                           CmdType = "Page.processNavigation"
	PageGetLayoutMetrics                            CmdType = "Page.getLayoutMetrics"
	ProfilerEnable                                  CmdType = "Profiler.enable"
	ProfilerDisable                                 CmdType = "Profiler.disable"
	ProfilerSetSamplingInterval                     CmdType = "Profiler.setSamplingInterval"
	ProfilerStart                                   CmdType = "Profiler.start"
	ProfilerStop                                    CmdType = "Profiler.stop"
	ProfilerStartPreciseCoverage                    CmdType = "Profiler.startPreciseCoverage"
	ProfilerStopPreciseCoverage                     CmdType = "Profiler.stopPreciseCoverage"
	ProfilerTakePreciseCoverage                     CmdType = "Profiler.takePreciseCoverage"
	ProfilerGetBestEffortCoverage                   CmdType = "Profiler.getBestEffortCoverage"
	RenderingSetShowPaintRects                      CmdType = "Rendering.setShowPaintRects"
	RenderingSetShowDebugBorders                    CmdType = "Rendering.setShowDebugBorders"
	RenderingSetShowFPSCounter                      CmdType = "Rendering.setShowFPSCounter"
	RenderingSetShowScrollBottleneckRects           CmdType = "Rendering.setShowScrollBottleneckRects"
	RenderingSetShowViewportSizeOnResize            CmdType = "Rendering.setShowViewportSizeOnResize"
	RuntimeEvaluate                                 CmdType = "Runtime.evaluate"
	RuntimeAwaitPromise                             CmdType = "Runtime.awaitPromise"
	RuntimeCallFunctionOn                           CmdType = "Runtime.callFunctionOn"
	RuntimeGetProperties                            CmdType = "Runtime.getProperties"
	RuntimeReleaseObject                            CmdType = "Runtime.releaseObject"
	RuntimeReleaseObjectGroup                       CmdType = "Runtime.releaseObjectGroup"
	RuntimeRunIfWaitingForDebugger                  CmdType = "Runtime.runIfWaitingForDebugger"
	RuntimeEnable                                   CmdType = "Runtime.enable"
	RuntimeDisable                                  CmdType = "Runtime.disable"
	RuntimeDiscardConsoleEntries                    CmdType = "Runtime.discardConsoleEntries"
	RuntimeSetCustomObjectFormatterEnabled          CmdType = "Runtime.setCustomObjectFormatterEnabled"
	RuntimeCompileScript                            CmdType = "Runtime.compileScript"
	RuntimeRunScript                                CmdType = "Runtime.runScript"
	SchemaGetDomains                                CmdType = "Schema.getDomains"
	SecurityEnable                                  CmdType = "Security.enable"
	SecurityDisable                                 CmdType = "Security.disable"
	SecurityShowCertificateViewer                   CmdType = "Security.showCertificateViewer"
	SecurityHandleCertificateError                  CmdType = "Security.handleCertificateError"
	SecuritySetOverrideCertificateErrors            CmdType = "Security.setOverrideCertificateErrors"
	ServiceWorkerEnable                             CmdType = "ServiceWorker.enable"
	ServiceWorkerDisable                            CmdType = "ServiceWorker.disable"
	ServiceWorkerUnregister                         CmdType = "ServiceWorker.unregister"
	ServiceWorkerUpdateRegistration                 CmdType = "ServiceWorker.updateRegistration"
	ServiceWorkerStartWorker                        CmdType = "ServiceWorker.startWorker"
	ServiceWorkerSkipWaiting                        CmdType = "ServiceWorker.skipWaiting"
	ServiceWorkerStopWorker                         CmdType = "ServiceWorker.stopWorker"
	ServiceWorkerInspectWorker                      CmdType = "ServiceWorker.inspectWorker"
	ServiceWorkerSetForceUpdateOnPageLoad           CmdType = "ServiceWorker.setForceUpdateOnPageLoad"
	ServiceWorkerDeliverPushMessage                 CmdType = "ServiceWorker.deliverPushMessage"
	ServiceWorkerDispatchSyncEvent                  CmdType = "ServiceWorker.dispatchSyncEvent"
	StorageClearDataForOrigin                       CmdType = "Storage.clearDataForOrigin"
	SystemInfoGetInfo                               CmdType = "SystemInfo.getInfo"
	TargetSetDiscoverTargets                        CmdType = "Target.setDiscoverTargets"
	TargetSetAutoAttach                             CmdType = "Target.setAutoAttach"
	TargetSetAttachToFrames                         CmdType = "Target.setAttachToFrames"
	TargetSetRemoteLocations                        CmdType = "Target.setRemoteLocations"
	TargetSendMessageToTarget                       CmdType = "Target.sendMessageToTarget"
	TargetGetTargetInfo                             CmdType = "Target.getTargetInfo"
	TargetActivateTarget                            CmdType = "Target.activateTarget"
	TargetCloseTarget                               CmdType = "Target.closeTarget"
	TargetAttachToTarget                            CmdType = "Target.attachToTarget"
	TargetDetachFromTarget                          CmdType = "Target.detachFromTarget"
	TargetCreateBrowserContext                      CmdType = "Target.createBrowserContext"
	TargetDisposeBrowserContext                     CmdType = "Target.disposeBrowserContext"
	TargetCreateTarget                              CmdType = "Target.createTarget"
	TargetGetTargets                                CmdType = "Target.getTargets"
	TetheringBind                                   CmdType = "Tethering.bind"
	TetheringUnbind                                 CmdType = "Tethering.unbind"
	TracingStart                                    CmdType = "Tracing.start"
	TracingEnd                                      CmdType = "Tracing.end"
	TracingGetCategories                            CmdType = "Tracing.getCategories"
	TracingRequestMemoryDump                        CmdType = "Tracing.requestMemoryDump"
	TracingRecordClockSyncMarker                    CmdType = "Tracing.recordClockSyncMarker"
)
