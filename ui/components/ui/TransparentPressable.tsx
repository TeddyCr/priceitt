import React from "react";
import { Pressable, StyleSheet } from "react-native";
import { styles } from "../styles/Generic";

export function TransparentPressable({children}: {children: React.ReactNode}) {
    return (
        <Pressable style={[sheetStyles.container]}>
            {children}
        </Pressable>
    )
}

const sheetStyles = StyleSheet.create({
    container: {
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'transparent',
        margin: 10,
        marginHorizontal: 15,
        borderRadius: 10,
        borderWidth: 1,
        borderColor: '#1b3c4b',
        height: 50,
    }
})