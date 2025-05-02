import React from 'react';
import { TouchableOpacity, Text, StyleSheet, View, Image } from 'react-native';
import { styles as genericStyles } from '@/components/styles/Generic';

interface SignInWithGoogleProps {
  onPress: () => void;
  text?: string;
}

const SignInWithGoogle: React.FC<SignInWithGoogleProps> = ({ 
  onPress, 
  text = 'Sign in with Google' 
}) => {
  return (
    <TouchableOpacity 
      style={styles.button} 
      onPress={onPress}
      activeOpacity={0.8}
    >
      <View style={styles.buttonContent}>
        {/* <Image 
          source={require('@/assets/images/google-logo.png')} 
          style={styles.logo}
          resizeMode="contain"
        /> */}
        <Text style={styles.text}>{text}</Text>
      </View>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    backgroundColor: 'white',
    borderRadius: 4,
    padding: 12,
    alignItems: 'center',
    justifyContent: 'center',
    marginVertical: 10,
    borderWidth: 1,
    borderColor: '#ddd',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 1,
    elevation: 2,
  },
  buttonContent: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
  },
  logo: {
    width: 24,
    height: 24,
    marginRight: 10,
  },
  text: {
    color: '#757575',
    fontSize: 16,
    fontWeight: '500',
  }
});

export default SignInWithGoogle;
