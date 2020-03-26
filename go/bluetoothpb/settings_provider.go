package bluetoothpb

type SettingsProvider interface {
	GetSettings() *Settings
}
