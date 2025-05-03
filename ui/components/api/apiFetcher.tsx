const baseUrl = process.env.EXPO_PUBLIC_BASE_API_URL;

export async function ApiFetch(endpoint: string, options: RequestInit) {
  const response = await fetch(`${baseUrl}${endpoint}`, {
    ...options,
    headers: {
      ...options.headers,
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    const error = new Error("Failed to fetch data");
    error.cause = await response.json();
    throw error;
  }

  return response.json();
}
