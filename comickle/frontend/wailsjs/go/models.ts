export namespace main {
	
	export class Comic {
	    id: number;
	    title: string;
	    path: string;
	    coverImage: string;
	
	    static createFrom(source: any = {}) {
	        return new Comic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.path = source["path"];
	        this.coverImage = source["coverImage"];
	    }
	}
	export class ComicPage {
	    index: number;
	    image: string;
	
	    static createFrom(source: any = {}) {
	        return new ComicPage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.image = source["image"];
	    }
	}

}

export namespace zip {
	
	export class File {
	    Name: string;
	    Comment: string;
	    NonUTF8: boolean;
	    CreatorVersion: number;
	    ReaderVersion: number;
	    Flags: number;
	    Method: number;
	    // Go type: time
	    Modified: any;
	    ModifiedTime: number;
	    ModifiedDate: number;
	    CRC32: number;
	    CompressedSize: number;
	    UncompressedSize: number;
	    CompressedSize64: number;
	    UncompressedSize64: number;
	    Extra: number[];
	    ExternalAttrs: number;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Comment = source["Comment"];
	        this.NonUTF8 = source["NonUTF8"];
	        this.CreatorVersion = source["CreatorVersion"];
	        this.ReaderVersion = source["ReaderVersion"];
	        this.Flags = source["Flags"];
	        this.Method = source["Method"];
	        this.Modified = this.convertValues(source["Modified"], null);
	        this.ModifiedTime = source["ModifiedTime"];
	        this.ModifiedDate = source["ModifiedDate"];
	        this.CRC32 = source["CRC32"];
	        this.CompressedSize = source["CompressedSize"];
	        this.UncompressedSize = source["UncompressedSize"];
	        this.CompressedSize64 = source["CompressedSize64"];
	        this.UncompressedSize64 = source["UncompressedSize64"];
	        this.Extra = source["Extra"];
	        this.ExternalAttrs = source["ExternalAttrs"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

