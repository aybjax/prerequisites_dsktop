class DbConnection {
    static _instance: DbConnection|undefined

    static get instance() {
        if(!DbConnection._instance) {
            DbConnection._instance = new DbConnection
        }

        return DbConnection._instance
    }

    static readonly DB_URL: string = btoa('DB_URL')
    static readonly DB_PORT: string = btoa('DB_PORT')
    static readonly DB_USER: string = btoa('DB_USER')
    static readonly DB_PASSWORD: string = btoa('DB_PASSWORD')

    private _url: string;
    private _port: string;
    private _user: string;
    private _passwd: string;

    private constructor() {
        this._url = localStorage.getItem(DbConnection.DB_URL)
        this._port = localStorage.getItem(DbConnection.DB_PORT)
        this._user = localStorage.getItem(DbConnection.DB_USER)
        this._passwd = localStorage.getItem(DbConnection.DB_PASSWORD)
    }

    get url() {
        if(!this._url) return '';

        return atob(this._url)
    }
    set url(url: string) {
        const encoded = btoa(url ?? '')
        localStorage.setItem(DbConnection.DB_URL, encoded)
        this._url = encoded
    }

    get port() {
        if(!this._port) return '';

        return atob(this._port)
    }
    set port(port: string) {
        const encoded = btoa(port ?? '')
        localStorage.setItem(DbConnection.DB_PORT, encoded)
        this._port = encoded
    }

    get user() {
        if(!this._user) return '';

        return atob(this._user)
    }
    set user(user: string) {
        const encoded = btoa(user ?? '')
        localStorage.setItem(DbConnection.DB_USER, encoded)
        this._user = encoded
    }

    get passwd() {
        if(!this._passwd) return '';

        return atob(this._passwd)
    }
    set passwd(passwd: string) {
        const encoded = btoa(passwd ?? '')
        localStorage.setItem(DbConnection.DB_PASSWORD, encoded)
        this._passwd = encoded
    }
}

export {DbConnection}