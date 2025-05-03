import { CreateUser } from "@/models/generated/createUser";
import { GoogleSignin } from "@react-native-google-signin/google-signin";
import { router } from "expo-router";
import { ApiEndpoints } from "./Endpoints";
import { ApiFetch } from "./ApiFetcher";

export async function CreateUserGoogle(): Promise<{username: string | null, idToken: string | null} | never> {
    await GoogleSignin.hasPlayServices();
    const currentUser = GoogleSignin.getCurrentUser();
    var createUser: CreateUser
    if (currentUser) {
        createUser = {
            name: currentUser.user.name ?? "",
            email: currentUser.user.email ?? "",
            authType: "google",
            authMechanism: {
                type: "google",
                idToken: currentUser.idToken ?? "",
                audience: process.env.EXPO_PUBLIC_GOOGLE_CLIENT_ID ?? ""
            }
        }
    } else {
        createUser = await getUserFromGoogle()
    }
    await ApiFetch(ApiEndpoints.USER, {
        method: 'POST',
        body: JSON.stringify(createUser)
    })
    return {username: currentUser?.user.email ?? null, idToken: currentUser?.idToken ?? null}
}

async function getUserFromGoogle(): Promise<CreateUser> {
    const userInfo = await GoogleSignin.signIn();
    if (userInfo.type === "cancelled") {
        throw new Error("User cancelled the sign in process");
    }
    const { idToken } = await GoogleSignin.getTokens();
    return {
        name: userInfo.data.user.name ?? "",
        email: userInfo.data.user.email ?? "",
        authType: "google",
        authMechanism: {
            type: "google",
            idToken: idToken ?? "",
            audience: process.env.EXPO_PUBLIC_GOOGLE_CLIENT_ID ?? ""
        }
    }
}