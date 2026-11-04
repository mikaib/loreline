package loreline.lib;

import go.Error;
import go.Errors;
import go.Tuple;
import loreline.Loreline;

@:keep
class GoLibrary {

    static function main() {}

    public static function toError(e:Any):Error {
        return Errors._new(Std.string(e));
    }

    public function new() {}

    @:go.Export
    @:go.Tuple('script', 'err')
    public function parse(source:String, path:String, readFile:String -> String):Tuple<{ script:GoScript, err:Error }> {
        try {
            final script = path == null || path.length == 0
                ? Loreline.parse(source)
                : Loreline.parse(source, path, (importPath, callback) -> callback(readFile(importPath)));
            if (script == null) {
                final e = Loreline.lastError();
                return { script: null, err: Errors._new(e != null ? Std.string(e) : "Failed to parse script") };
            }
            return { script: new GoScript(script), err: null };
        }
        catch (e:Any) {
            return { script: null, err: toError(e) };
        }
    }

}
