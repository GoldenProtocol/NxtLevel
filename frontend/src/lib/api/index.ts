enum Method {
    POST = "POST",
    GET = "GET"
}

class Backend {
    private host: string;
    private method: Method | undefined;  
    private path: string | undefined; 
    constructor(host:string) {
        this.host = host; 
    }

    create() {
        this.method = Method.POST
        return this
    }

    read() {
        this.method = Method.GET
        return this
    }

    companies() {
        this.path = "/v1/companies"
        return this
    }

    nodes() {
        this.path = "/v1/nodes"
        return this
    }

    games() {
        this.path = "/v1/games"
        return this
    }

    connections() {
        this.path = "/v1/connections"
        return this
    }

    async send<T, R>(body: T | null) {
        const res = await fetch(`http://${this.host}${this.path}`, {
            method: this.method, 
            headers: {
                "Content-Type": "application/json"
            }, 
            body: body ? JSON.stringify(body): undefined
        }) 

        if (res.ok) {
            return (await res.json()) as R 
        }

        throw new Error("Bad Request")
    }
}

export default new Backend("localhost:8080")