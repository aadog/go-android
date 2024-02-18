package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

var Tools ToolsStruct

type ToolsStruct struct {
}

// GetApplicationContext request androidContext
func (t ToolsStruct) GetApplicationContext() mo.Result[*jvm.ObjectWrapper] {
	androidAppActivityClass := AndroidAppActivityThreadClass()
	currentActivityThread := androidAppActivityClass.CallStaticObjectA("currentActivityThread").MustGet()
	defer currentActivityThread.Free()
	application := currentActivityThread.CallObjectA("getApplication").MustGet()
	defer application.Free()
	applicationContext := application.CallObjectA("getApplicationContext").MustGet()
	return mo.Ok(applicationContext)
}
func (t ToolsStruct) GetPackageName(applicationContext *jvm.ObjectWrapper) mo.Result[string] {
	packageName := applicationContext.CallStringA("getPackageName").MustGet()
	return mo.Ok(packageName)
}
func (t ToolsStruct) GetVersionName(applicationContext *jvm.ObjectWrapper) mo.Result[string] {
	packageName := t.GetPackageName(applicationContext).MustGet()
	pkgManager := applicationContext.CallObjectA("getPackageManager").MustGet()
	defer pkgManager.Free()
	pkgInfo := pkgManager.CallObjectA("getPackageInfo", packageName, 0).MustGet()
	defer pkgInfo.Free()
	str := pkgInfo.GetStringFieldValue("versionName").MustGet()
	return mo.Ok(str)
}
