import { ApiEndpoints } from "./Endpoints";
import { ApiFetch } from "./ApiFetcher";

export async function CreateUserBasic(
  fullName: string,
  email: string,
  password: string,
  confirmPassword: string,
): Promise<void | never> {
  await ApiFetch(ApiEndpoints.USER, {
    method: "POST",
    body: JSON.stringify({
      name: fullName,
      email: email,
      authType: "basic",
      authMechanism: {
        type: "basic",
        password: password,
        confirmPassword: confirmPassword,
      },
    }),
  });
}
