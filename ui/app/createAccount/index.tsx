import { apiFetch } from "@/components/api/apiFetcher";
import { styles } from "@/components/styles/Generic";
import { TopNav } from "@/components/styles/TopNav";
import { Alert } from "@/components/ui/Alert";
import GreenPressable from "@/components/ui/GreenPressable";
import { TransparentTextInput } from "@/components/ui/TransparentTextInput";
import ValidateEmail from "@/components/validators/ValidateEmail";
import ValidatePassword, { ValidateConfirmPassword } from "@/components/validators/ValidatePassword";
import { router } from "expo-router";
import { useEffect, useState } from "react";
import { View, Text, SafeAreaView } from "react-native";

export default function CreateAccount() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [fullName, setFullName] = useState('');
    const [alert, setAlert] = useState<{type: 'error' | 'success' | 'info', message: string} | null>(null);
    
    const handleDismissAlert = () => {
        setAlert(null);
    }

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
        }
    }, [alert])


    return (
        <SafeAreaView
        style={styles.defaultBackgroundContainer}>
            <TopNav />
            {
                alert && (
                    <Alert
                        type={alert.type}
                        message={alert.message}
                        onDismiss={handleDismissAlert}
                    />
                )
            }
            <View style={{flex: 1}}>
                <Text style={styles.titleText}>Create Account</Text>
            </View>
            <TransparentTextInput
                placeholder="Full Name"
                text={fullName}
                onChangeText={setFullName}
                validator={(fullName) => {return null}}
            />
            <TransparentTextInput
                placeholder="Email"
                text={email}
                onChangeText={setEmail}
                validator={(email) => {return ValidateEmail(email)}}
            />
            <TransparentTextInput
                placeholder="Password"
                text={password}
                onChangeText={setPassword}
                secureTextEntry={true}
                validator={(password) => {return ValidatePassword(password)}}
                showPasswordToggle={true}
            />
            <TransparentTextInput
                placeholder="Confirm Password"
                text={confirmPassword}
                onChangeText={setConfirmPassword}
                secureTextEntry={true}
                validator={(confirmPassword) => {return ValidateConfirmPassword(password, confirmPassword)}}
                showPasswordToggle={true}
            />
            <GreenPressable 
                disabled={!email || !password || !confirmPassword || !fullName || 
                          ValidateEmail(email) !== null || 
                          ValidatePassword(password) !== null || 
                          ValidateConfirmPassword(password, confirmPassword) !== null}
                onPress={async () => {
                    try {
                        await apiFetch('/api/v1/user', {
                            method: 'POST',
                            body: JSON.stringify({
                                name: fullName,
                                email: email,
                                password: password,
                                confirmPassword: confirmPassword,
                                authType: 'basic'
                            })
                        })
                        router.push('/');
                    } catch (error) {
                        if (error instanceof Error && error.cause) {
                            const errorData = error.cause as { message: string };
                            setAlert({
                                type: 'error',
                                message: errorData.error,
                            });
                        } else {
                            setAlert({
                                type: 'error',
                                message: 'An unknown error occurred',
                            });
                        }
                    }
                }}
            >
            <Text style={styles.greenPressableText}>Create Account</Text>
            </GreenPressable>
            <View style={{flex: 1, justifyContent: 'top', alignItems: 'center', paddingTop: 25}}>
                <Text>Already have an account? <Text onPress={() => router.push('/login')} style={{color: '#2ca963'}}>Login</Text></Text>
            </View>
        </SafeAreaView>
    );
}