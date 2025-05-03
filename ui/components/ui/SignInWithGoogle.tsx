import React, { useEffect } from "react";
import { Text, StyleSheet, View, Image, Pressable } from "react-native";
import { GoogleSignin } from "@react-native-google-signin/google-signin";
import { styles } from "@/components/styles/Generic";

import { CreateUserGoogle } from "../api/CreateUserGoogle";

interface SignInWithGoogleProps {
  text?: string;
  onError: (message: string) => void;
  onSuccess: (email: string, idToken: string) => void;
}

export default function SignInWithGoogle({
  text = "Sign in with Google",
  onSuccess,
  onError,
}: SignInWithGoogleProps): React.ReactNode {
  useEffect(() => {
    GoogleSignin.configure({
      iosClientId: process.env.EXPO_PUBLIC_GOOGLE_CLIENT_ID,
    });
  }, []);

  const onPress = async () => {
    try {
      const { username, idToken } = await CreateUserGoogle();
      if (username && idToken) {
        onSuccess(username, idToken);
      }
    } catch (error: unknown) {
      if (error instanceof Error) {
        onError(error.message);
      } else {
        onError(String(error));
      }
    }
  };

  return (
    <View style={sheetStyles.button}>
      <Pressable style={sheetStyles.buttonContent} onPress={onPress}>
        <Image
          source={require("@/assets/images/google-logo.png")}
          style={sheetStyles.logo}
          resizeMode="contain"
        />
        <Text style={styles.text}>{text}</Text>
      </Pressable>
    </View>
  );
}

const sheetStyles = StyleSheet.create({
  button: {
    backgroundColor: "white",
    borderWidth: 1,
    borderColor: "#1b3c4b",
    height: 50,
    borderRadius: 10,
    padding: 10,
    margin: 8,
    marginHorizontal: 15,
    alignItems: "center",
    justifyContent: "center",
    marginVertical: 10,
    elevation: 5,
  },
  buttonContent: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
  },
  logo: {
    width: 24,
    height: 24,
    marginRight: 10,
  },
});
