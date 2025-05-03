export default function ValidatePassword(password: string) {
  if (!password) {
    return "Password is required";
  }

  if (password.length < 16) {
    return "Password must be at least 16 characters long";
  }

  if (!/[A-Z]/.test(password)) {
    return "Password must include at least one uppercase letter";
  }

  if (!/[a-z]/.test(password)) {
    return "Password must include at least one lowercase letter";
  }

  if (!/[0-9]/.test(password)) {
    return "Password must include at least one number";
  }

  if (!/[!@#$%^&*]/.test(password)) {
    return "Password must include at least one special character (!@#$%^&*)";
  }

  return null;
}

export function ValidateConfirmPassword(
  password: string,
  confirmPassword: string,
) {
  if (!confirmPassword) {
    return "Confirm password is required";
  }

  if (password !== confirmPassword) {
    return "Passwords do not match";
  }

  return null;
}
