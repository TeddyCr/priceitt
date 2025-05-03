export default function ValidateEmail(email: string) {
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

  if (!email) {
    return "Email is required";
  }
  if (!emailRegex.test(email)) {
    return "Invalid email address";
  }

  return null;
}
