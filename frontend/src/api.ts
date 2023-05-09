const ADDRESS = "http://localhost"
const PORT = "3333"
const url = `${ADDRESS}:${PORT}/watchlist`

async function request<TResponse>(
    config: RequestInit = {}
): Promise<TResponse> {
    const response = await fetch(url, config)
    return await response.json().then(r => r)
}

const api = {
    get: async <TResponse>(config: RequestInit = {method:"GET", mode:"cors"}) => await request<TResponse>(config),
    post: async <TBody extends BodyInit, TResponse>(body: TBody) =>
        await request<TResponse>({method: 'POST', body}),
}

export default api