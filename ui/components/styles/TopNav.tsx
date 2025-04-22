import { Link } from "expo-router";
import { View } from "react-native";
import { styles } from "./Generic";
import { StyleSheet } from "react-native";
import { Image } from "expo-image";

const unstackedLogo = require('../../assets/images/unstacked_logo.png');

export function TopNav() {
    return (
    <View style={fnStyles.topNav}>
        <Image source={unstackedLogo}
            style={{
            flex: 4,
            justifyContent: 'center',
            height: 40,
            }}
            contentFit="scale-down"
        />
    </View>
    )
}

const fnStyles = StyleSheet.create({
    topNav: {
        paddingRight: 30,
        paddingLeft: 0,
        flex: 1,
        flexDirection: 'row'
    }
})