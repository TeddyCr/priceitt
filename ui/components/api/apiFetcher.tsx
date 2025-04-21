const baseUrl = __DEV__ ? 'http://localhost:8000' : 'https://priceitt.xyz';

export async function apiFetch(endpoint: string, options: RequestInit) {
    const response = await fetch(
        `${baseUrl}${endpoint}`,
        {
            ...options,
            headers: {
                ...options.headers,
                'Content-Type': 'application/json',
            },
        }
    )

    if (!response.ok) {
        const error = new Error('Failed to fetch data');
        error.cause = await response.json();
        throw error;
    }

    return response.json();
}