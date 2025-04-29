import { styles } from "@/components/styles/Generic";
import { TransparentPressable } from "@/components/ui/TransparentPressable";
import GreenPressable from "@/components/ui/GreenPressable";
import {useAuthSession} from "@/providers/AuthProvider";
import { Image } from 'expo-image';
import Uuid from "expo-modules-core/src/uuid";
import { Link, router } from "expo-router";
import {ReactNode, useEffect} from "react";
import {Text, View} from "react-native";
import { TopNav } from "@/components/styles/TopNav";
import { SafeAreaView } from "react-native-safe-area-context";
import SignInWithGoogle from "@/components/ui/SignInWithGoogle";
import { GoogleSignin } from "@react-native-google-signin/google-signin";

export default function Login(): ReactNode {
  const {signIn} = useAuthSession();
  const login = ():void => {
    const random: string = Uuid.v4();
    signIn(random);
  }

  useEffect(() => {
    GoogleSignin.configure({
      iosClientId: "713461610853-longvrf745bi44vl5p7k96rashc46pan.apps.googleusercontent.com"
    });
  }, []);

  const signInWithGoogle = async () => {
    try {
      await GoogleSignin.hasPlayServices();
      const userInfo = await GoogleSignin.signIn();
      const { accessToken, idToken } = await GoogleSignin.getTokens();
      console.log(userInfo);
      console.log(accessToken);
      console.log(idToken);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <SafeAreaView
      style={styles.defaultBackgroundContainer}
    >
      <TopNav />
      <View style={{flex: 2}}>
        <Text style={styles.titleText}>Save On Groceries</Text>
      </View>
      <Image source={require('../../assets/images/shopper.png')}
        style={{
          paddingTop: 30,
          marginBottom: 10,
          flex: 6,
        }}
        contentFit="cover"
      />
      <View style={{flex: 2}}>
        <GreenPressable onPress={() => { router.push('/createAccount') }}>
          <Text style={styles.greenPressableText}>Create Account</Text>
        </GreenPressable>
        <TransparentPressable>
          <Text style={styles.transparentPressableText}>Login</Text>
        </TransparentPressable>
        <SignInWithGoogle onPress={signInWithGoogle} />
      </View>
    </SafeAreaView>
  );
}