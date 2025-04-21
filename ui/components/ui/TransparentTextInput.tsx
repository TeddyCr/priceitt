import { useState } from "react";
import { TextInput, StyleSheet, Text, View, Pressable } from "react-native";
import { styles } from "../styles/Generic";
import { Ionicons } from "@expo/vector-icons";

type TransparentFormInputProps = {
    placeholder: string;
    text: string;
    onChangeText: (text: string) => void;
    secureTextEntry?: boolean;
    validator: (text: string) => string | null;
    showPasswordToggle?: boolean;
}

export function TransparentTextInput({
    placeholder, 
    text, 
    onChangeText, 
    secureTextEntry, 
    validator,
    showPasswordToggle = false
}: TransparentFormInputProps) {
    const [error, setError] = useState<string | null>(null);
    const [showPassword, setShowPassowrd] = useState(false);

    const handleBlur = () => {
        const error = validator(text);
        setError(error);
    }
    
    return (
        <View>
            <TextInput
                style={styleSheet.formInput}
                placeholder={placeholder}
                value={text}
                onChangeText={t => {if (error) setError(null); onChangeText(t)}} // if there is an error, clear the error
                placeholderTextColor="#1b3c4b"
                secureTextEntry={secureTextEntry && !showPassword}
                onBlur={handleBlur}
                autoCapitalize="none"
        />
        {
            showPasswordToggle && (
                <Pressable 
                    style={styleSheet.eyeIcon}
                    onPress={() => setShowPassowrd(!showPassword)}
                >
                    <Ionicons 
                        name={showPassword ? "eye-outline" : "eye-off-outline"} 
                        size={24} 
                        color="#1b3c4b" 
                    />
                </Pressable>
            )
        }
            {error && <Text style={styles.errorText}>{error}</Text>}
        </View>
    )
}

const styleSheet = StyleSheet.create({
    formInput: {
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'transparent',
        padding: 10,
        borderRadius: 10,
        borderWidth: 2,
        borderColor: '#1b3c4b',
        height: 50,
        margin: 8,
        marginHorizontal: 15,
    },
    eyeIcon: {
        position: 'absolute',
        right: 25,
        padding: 10,
        top: 10,
    }
})