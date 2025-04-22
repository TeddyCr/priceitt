import React from "react";
import { Pressable, StyleSheet } from "react-native";

export function TransparentPressable({children}: {children: React.ReactNode}) {
    return (
        <Pressable style={styles.container}>
            {children}
        </Pressable>
    )
}

const styles = StyleSheet.create({
    container: {
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'transparent',
        margin: 10,
        marginHorizontal: 15,
        borderRadius: 10,
        borderWidth: 2,
        borderColor: '#1b3c4b',
        height: 50,
    }
})