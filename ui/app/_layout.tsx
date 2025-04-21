import { Slot, Stack } from 'expo-router';
import 'react-native-reanimated';

import AuthProvider from '@/providers/AuthProvider';
import { ReactNode } from 'react';

export default function RootLayout(): ReactNode {
  return (
    <AuthProvider>
      <Slot />
    </AuthProvider>
  );
}
