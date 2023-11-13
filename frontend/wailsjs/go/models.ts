export namespace main {
	
	export class Conn {
	    uri: string;
	    port: string;
	    user: string;
	    pswd: string;
	
	    static createFrom(source: any = {}) {
	        return new Conn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uri = source["uri"];
	        this.port = source["port"];
	        this.user = source["user"];
	        this.pswd = source["pswd"];
	    }
	}

}

export namespace repo {
	
	export class Edge {
	    id: string;
	    weight: number;
	
	    static createFrom(source: any = {}) {
	        return new Edge(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.weight = source["weight"];
	    }
	}

}

