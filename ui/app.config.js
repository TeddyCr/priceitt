import 'dotenv/config'; // Optional: Load environment variables from .env

export default () => ({
  expo: {
    name: "ui",
    slug: "ui",
    version: "1.0.0",
    orientation: "portrait",
    icon: "./assets/images/icon.png",
    scheme: "PriceItt",
    userInterfaceStyle: "automatic",
    newArchEnabled: true,

    ios: {
      supportsTablet: true,
      bundleIdentifier: "com.teddycr.ui",
      infoPlist: {
        ITSAppUsesNonExemptEncryption: false
      }
    },

    android: {
      package: "com.teddycr.ui"
    },

    web: {
      bundler: "metro",
      output: "static",
      favicon: "./assets/images/favicon.png"
    },

    plugins: [
      "expo-router",
      [
        "expo-splash-screen",
        {
          image: "./assets/images/splash-icon.png",
          imageWidth: 200,
          resizeMode: "contain",
          backgroundColor: "#ffffff"
        }
      ],
      [
        "@react-native-google-signin/google-signin",
        {
          iosUrlScheme: process.env.EXPO_PUBLIC_GOOGLE_URL_SCHEME
        }
      ],
      "expo-font"
    ],

    experiments: {
      typedRoutes: true
    },

    extra: {
      router: {
        origin: false
      },
      eas: {
        projectId: "1aff4574-f08e-4dbc-9e02-f0b14850d0df"
      },
      // Add dynamic env values here if needed
      apiUrl: process.env.API_URL
    }
  }
});
