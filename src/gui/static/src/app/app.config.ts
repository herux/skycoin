export const AppConfig = {
  otcEnabled: false,
  maxHardwareWalletAddresses: 1,
  useHwWalletDaemon: true,

  urlForVersionChecking: 'https://version.skycoin.net/skycoin/version.txt',
  walletDownloadUrl: 'https://www.skycoin.net/downloads/',

  languages: [{
      code: 'en',
      name: 'English',
      iconName: 'en.png',
    },
    {
      code: 'zh',
      name: '中文',
      iconName: 'zh.png',
    },
    {
      code: 'es',
      name: 'Español',
      iconName: 'es.png',
    },
  ],
  defaultLanguage: 'en',
};
