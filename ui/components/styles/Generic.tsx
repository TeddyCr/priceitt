import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
  transparentPressableText: {
    fontSize: 18,
    color: "#1b3c4b",
  },
  text: {
    color: "#1b3c4b",
    fontSize: 16,
  },
  greenPressableText: {
    color: "#faf8f5",
    fontSize: 18,
  },
  defaultBackgroundContainer: {
    backgroundColor: "#faf8f5",
    flex: 1,
    flexDirection: "column",
    paddingTop: 20,
    paddingBottom: 20,
  },
  titleText: {
    paddingTop: 20,
    textAlign: "center",
    justifyContent: "center",
    fontSize: 48,
    fontWeight: "bold",
    flexWrap: "wrap",
    width: "100%",
    color: "#1b3c4b",
    lineHeight: 45,
  },
  topNav: {
    textAlign: "right",
    verticalAlign: "middle",
    fontWeight: "bold",
    fontSize: 16,
  },
  errorText: {
    color: "#f00",
    marginBottom: 4,
    fontSize: 12,
    marginHorizontal: 15,
  },
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "transparent",
    margin: 10,
  },
});
