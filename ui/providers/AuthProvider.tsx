import AsyncStorage from "@react-native-async-storage/async-storage";
import * as SecureStore from 'expo-secure-store';
import { router } from "expo-router";
import { createContext, MutableRefObject, useCallback, useContext, useEffect, useRef, useState } from "react";

const AuthContext = createContext<{
    signIn: (accessToken: string, refreshToken: string) => void;
    signOut: () => void;
    accessToken: MutableRefObject<string | null> | null;
    refreshToken: MutableRefObject<string | null> | null;
    isLoading: boolean;
}>({
    signIn: () => null,
    signOut: () => null,
    accessToken: null,
    refreshToken: null,
    isLoading: true
});

export function useAuthSession() {
    return useContext(AuthContext);
}

export default function AuthProvider({children}: {children: React.ReactNode}) {
    const accessTokenRef = useRef<string | null>(null);
    const refreshTokenRef = useRef<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        (async ():Promise<void> => {
            const accessToken = await SecureStore.getItemAsync('accessToken');
            const refreshToken = await SecureStore.getItemAsync('refreshToken');
            accessTokenRef.current = accessToken || null;
            refreshTokenRef.current = refreshToken || null;
            setIsLoading(false);
        })()
    }, []);

    const signIn = useCallback(async (accessToken: string, refreshToken: string) => {
        await SecureStore.setItemAsync('accessToken', accessToken);
        await SecureStore.setItemAsync('refreshToken', refreshToken);
        accessTokenRef.current = accessToken;
        refreshTokenRef.current = refreshToken;
        router.replace('/');
    }, [])

    const signOut = useCallback(async () => {
        await SecureStore.deleteItemAsync('accessToken');
        await SecureStore.deleteItemAsync('refreshToken');
        accessTokenRef.current = null;
        refreshTokenRef.current = null;
        router.replace('/');
    }, [])

    return (
        <AuthContext.Provider
            value={{
                signIn,
                signOut,
                accessToken: accessTokenRef,
                refreshToken: refreshTokenRef,
                isLoading,
            }}
        >
            {children}
        </AuthContext.Provider>
    )
}
