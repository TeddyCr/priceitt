import { styles } from "@/components/styles/Generic";
import { TransparentPressable } from "@/components/ui/TransparentPressable";
import GreenPressable from "@/components/ui/GreenPressable";
import { Image } from "expo-image";
import { router } from "expo-router";
import { ReactNode } from "react";
import { Text, View } from "react-native";
import { TopNav } from "@/components/styles/TopNav";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Login(): ReactNode {
  return (
    <SafeAreaView style={styles.defaultBackgroundContainer}>
      <TopNav />
      <View style={{ flex: 2 }}>
        <Text style={styles.titleText}>Save On Groceries</Text>
      </View>
      <Image
        source={require("../../assets/images/shopper.png")}
        style={{
          paddingTop: 30,
          marginBottom: 10,
          flex: 6,
        }}
        contentFit="cover"
      />
      <View style={{ flex: 2 }}>
        <GreenPressable
          onPress={() => {
            router.push("/createAccount");
          }}
        >
          <Text style={styles.greenPressableText}>Create Account</Text>
        </GreenPressable>
        <TransparentPressable>
          <Text style={styles.transparentPressableText}>Login</Text>
        </TransparentPressable>
      </View>
    </SafeAreaView>
  );
}
