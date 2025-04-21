import React from "react";
import { Pressable, StyleSheet } from "react-native";

type GreenPressableProps = {
    children: React.ReactNode;
    onPress?: () => void;
    disabled?: boolean;
}

const GreenPressable: React.FC<GreenPressableProps> = ({children, onPress = () => {}, disabled = false}) => {
    return (
        <Pressable 
            style={[
                styles.container, 
                disabled && styles.disabledContainer
            ]} 
            onPress={onPress} 
            disabled={disabled}>
            {children}
        </Pressable>
    )
}

export default GreenPressable;

const styles = StyleSheet.create({
    container: {
        height: 50,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#2ca963',
        margin: 5,
        marginHorizontal: 15,
        borderRadius: 10,
        borderWidth: 2,
        borderColor: '#2ca963',
    },
    disabledContainer: {
        backgroundColor: '#a5d8bc', // lighter shade of #2ca963
        borderColor: '#a5d8bc',     // lighter shade of #2ca963
    }
})