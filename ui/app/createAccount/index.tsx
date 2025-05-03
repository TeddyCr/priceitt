import { CreateUserBasic } from "@/components/api/CreateUserBasic";
import { LoginUserBasic, LoginUserGoogle } from "@/components/api/LoginUser";
import { styles } from "@/components/styles/Generic";
import { TopNav } from "@/components/styles/TopNav";
import { Alert } from "@/components/ui/Alert";
import GreenPressable from "@/components/ui/GreenPressable";
import SignInWithGoogle from "@/components/ui/SignInWithGoogle";
import { TransparentTextInput } from "@/components/ui/TransparentTextInput";
import ValidateEmail from "@/components/validators/ValidateEmail";
import ValidatePassword, {
  ValidateConfirmPassword,
} from "@/components/validators/ValidatePassword";
import { useAuthSession } from "@/providers/AuthProvider";
import { router } from "expo-router";
import { useEffect, useState } from "react";
import { View, Text, SafeAreaView, StyleSheet } from "react-native";

export default function CreateAccount() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [fullName, setFullName] = useState("");
  const [alert, setAlert] = useState<{
    type: "error" | "success" | "info";
    message: string;
  } | null>(null);

  const { signIn } = useAuthSession();
  const login = (accessToken: string, refreshToken: string) => {
    signIn(accessToken, refreshToken);
  };

  const handleDismissAlert = () => {
    setAlert(null);
  };

  const handleOnError = (message: string) => {
    setAlert({
      type: "error",
      message: message,
    });
  };

  const handleOnSuccess = async (email: string, idToken: string) => {
    const { accessToken, refreshToken } = await LoginUserGoogle(email, idToken);
    login(accessToken, refreshToken);
    router.push("/login");
  };

  useEffect(() => {
    let timer: NodeJS.Timeout;

    if (alert) {
      timer = setTimeout(() => {
        setAlert(null);
      }, 5000);
    }

    return () => {
      if (timer) {
        clearTimeout(timer);
      }
    };
  }, [alert]);

  return (
    <SafeAreaView style={styles.defaultBackgroundContainer}>
      <TopNav />
      {alert && (
        <Alert
          type={alert.type}
          message={alert.message}
          onDismiss={handleDismissAlert}
        />
      )}
      <View style={{ flex: 1, marginBottom: 40 }}>
        <Text style={styles.titleText}>Create Account</Text>
      </View>
      <SignInWithGoogle
        text="Continue with Google"
        onError={handleOnError}
        onSuccess={handleOnSuccess}
      />
      <View style={sheetStyles.orContainer}>
        <View style={sheetStyles.horizontalLine} />
        <Text style={sheetStyles.orText}>OR</Text>
        <View style={sheetStyles.horizontalLine} />
      </View>
      <TransparentTextInput
        placeholder="Full Name"
        text={fullName}
        onChangeText={setFullName}
        validator={(fullName) => {
          return null;
        }}
      />
      <TransparentTextInput
        placeholder="Email"
        text={email}
        onChangeText={setEmail}
        validator={(email) => {
          return ValidateEmail(email);
        }}
      />
      <TransparentTextInput
        placeholder="Password"
        text={password}
        onChangeText={setPassword}
        secureTextEntry={true}
        validator={(password) => {
          return ValidatePassword(password);
        }}
        showPasswordToggle={true}
      />
      <TransparentTextInput
        placeholder="Confirm Password"
        text={confirmPassword}
        onChangeText={setConfirmPassword}
        secureTextEntry={true}
        validator={(confirmPassword) => {
          return ValidateConfirmPassword(password, confirmPassword);
        }}
        showPasswordToggle={true}
      />
      <GreenPressable
        disabled={
          !email ||
          !password ||
          !confirmPassword ||
          !fullName ||
          ValidateEmail(email) !== null ||
          ValidatePassword(password) !== null ||
          ValidateConfirmPassword(password, confirmPassword) !== null
        }
        onPress={async () => {
          try {
            await CreateUserBasic(fullName, email, password, confirmPassword);
            const { accessToken, refreshToken } = await LoginUserBasic(
              email,
              password,
            );
            login(accessToken, refreshToken);
            router.push("/login");
          } catch (error) {
            if (error instanceof Error && error.cause) {
              const errorData = error.cause as {
                message: string;
                error: string;
              };
              setAlert({
                type: "error",
                message: errorData.error,
              });
            } else if (error instanceof Error) {
              setAlert({
                type: "error",
                message: error.message,
              });
            } else {
              setAlert({
                type: "error",
                message: String(error),
              });
            }
          }
        }}
      >
        <Text style={styles.greenPressableText}>Create Account</Text>
      </GreenPressable>
      <View
        style={{
          flex: 1,
          justifyContent: "top",
          alignItems: "center",
          paddingTop: 25,
        }}
      >
        <Text>
          Already have an account?{" "}
          <Text
            onPress={() => router.push("/login")}
            style={{ color: "#2ca963" }}
          >
            Login
          </Text>
        </Text>
      </View>
    </SafeAreaView>
  );
}

const sheetStyles = StyleSheet.create({
  orContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    marginVertical: 15,
    width: "100%",
  },
  horizontalLine: {
    flex: 1,
    height: 2,
    backgroundColor: "#2c4d5c",
    marginHorizontal: 15,
  },
  orText: {
    color: "#2c4d5c",
    fontSize: 16,
  },
});
