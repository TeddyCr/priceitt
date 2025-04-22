import AsyncStorage from "@react-native-async-storage/async-storage";
import { isLoading } from "expo-font";
import { router } from "expo-router";
import { createContext, MutableRefObject, useCallback, useContext, useEffect, useRef, useState } from "react";

const AuthContext = createContext<{
    signIn: (arg0: string) => void;
    signOut: () => void;
    token: MutableRefObject<string | null> | null;
    isLoading: boolean;
}>({
    signIn: () => null,
    signOut: () => null,
    token: null,
    isLoading: true
});

export function useAuthSession() {
    return useContext(AuthContext);
}

export default function AuthProvider({children}: {children: React.ReactNode}) {
    const tokenRef = useRef<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        (async ():Promise<void> => {
            AsyncStorage.clear()
            const token = await AsyncStorage.getItem('@accessToken');
            tokenRef.current = token || null;
            setIsLoading(false);
        })()
    }, []);

    const signIn = useCallback(async (token: string) => {
        await AsyncStorage.setItem('@accessToken', token);
        tokenRef.current = token;
        router.replace('/');
    }, [])

    const signOut = useCallback(async () => {
        await AsyncStorage.removeItem('@accessToken');
        tokenRef.current = null;
        router.replace('/login');
    }, [])

    return (
        <AuthContext.Provider
            value={{
                signIn,
                signOut,
                token: tokenRef,
                isLoading,
            }}
        >
            {children}
        </AuthContext.Provider>
    )
}
