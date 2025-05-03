import React from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { Ionicons } from "@expo/vector-icons";

export type AlertType = "success" | "error" | "info";

interface AlertProps {
  type: AlertType;
  message: string;
  onDismiss: () => void;
}

const getAlertColor = (type: AlertType) => {
  switch (type) {
    case "success":
      return "#2ca963"; // Using your app's green color
    case "error":
      return "#dc3545";
    case "info":
      return "#1b3c4b"; // Using your app's dark color
    default:
      return "#1b3c4b";
  }
};

const getIconName = (type: AlertType) => {
  switch (type) {
    case "success":
      return "checkmark-circle";
    case "error":
      return "alert-circle";
    case "info":
      return "information-circle";
    default:
      return "information-circle";
  }
};

export function Alert({ type, message, onDismiss }: AlertProps) {
  const alertColor = getAlertColor(type);
  const iconName = getIconName(type);

  return (
    <View style={[styles.container, { borderColor: alertColor }]}>
      <View style={styles.contentContainer}>
        <Ionicons
          name={iconName}
          size={24}
          color={alertColor}
          style={styles.icon}
        />
        <Text style={[styles.message, { color: alertColor }]}>{message}</Text>
      </View>
      <Pressable onPress={onDismiss} style={styles.dismissButton}>
        <Ionicons name="close" size={24} color={alertColor} />
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    padding: 12,
    marginHorizontal: 15,
    marginVertical: 8,
    borderWidth: 1,
    borderRadius: 8,
    backgroundColor: "#faf8f5", // Using your app's background color
  },
  contentContainer: {
    flexDirection: "row",
    alignItems: "center",
    flex: 1,
  },
  icon: {
    marginRight: 8,
  },
  message: {
    fontSize: 16,
    flex: 1,
  },
  dismissButton: {
    padding: 4,
  },
});
