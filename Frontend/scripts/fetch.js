class Fetch {
        static async get(url) {
                const response = await fetch(url, {
                        method: 'GET',
                        headers: {
                                'Content-Type': 'application/json'
                        }
                });
                return response.json()
        }
        static async post(url, body = null) {
                const response = await fetch(url, {
                        method: 'POST',
                        headers: {
                                'Content-Type': 'application/json',
                        },
                        body: (typeof body == 'object') && JSON.stringify(body)
                });

                return response.json()
        }
}
