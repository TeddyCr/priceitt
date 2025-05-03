import { useAuthSession } from "@/providers/AuthProvider";
import { ApiFetch } from "./ApiFetcher";
import { ApiEndpoints } from "./Endpoints";
import { Token } from "@/models/generated/token";

export async function LoginUserBasic(username: string, password: string): Promise<{accessToken: string, refreshToken: string} | never> {
    const response = await ApiFetch(ApiEndpoints.LOGIN, {
        method: 'POST',
        body: JSON.stringify({
            type: 'basic',
            username: username,
            data: {
                password: password
            }
        })
    })

    const data = response;
    const accessToken: Token = data.access.token;
    const refreshToken: Token = data.refresh.token;
    return {accessToken: accessToken.token, refreshToken: refreshToken.token}; 
}

export async function LoginUserGoogle(username: string, idToken: string): Promise<{accessToken: string, refreshToken: string} | never> {
    const response = await ApiFetch(ApiEndpoints.LOGIN, {
        method: 'POST',
        body: JSON.stringify({
            type: 'google',
            username: username,
            data: {
                idToken: idToken
            }
        })
    })

    const data = response;
    const accessToken: Token = data.access.token;
    const refreshToken: Token = data.refresh.token;
    return {accessToken: accessToken.token, refreshToken: refreshToken.token}; 
}