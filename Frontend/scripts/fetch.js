class Fetch {
        static async get(url, options = 'application/json') {
                const response = await fetch(url, {
                        method: 'GET',
                        headers: {
                                'Content-Type': options
                        }
                });
                return response.json()
        }
        static async post(url, options = 'application/json', body = null) {
                const response = await fetch(url, {
                        method: 'POST',
                        mode: 'no-cors',
                        headers: {
                                'Content-Type': options,
                        },
                        body: (typeof body == 'object') && JSON.stringify(body)
                });
                return response.json()
        }
}
