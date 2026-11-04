package loreline.lib;

import go.Error;
import go.Tuple;
import haxe.Json;
import loreline.Interpreter;

class GoInterpreter {

    public var interpreter:Interpreter;

    public function new() {}

    public static function beatOrNull(beat:String):String {
        return beat == null || beat.length == 0 ? null : beat;
    }

    @:go.Export
    public function start(beat:String):Error {
        try {
            interpreter.start(beatOrNull(beat));
            return null;
        }
        catch (e:Any) {
            return GoLibrary.toError(e);
        }
    }

    @:go.Export
    @:go.Tuple('save', 'err')
    public function save():Tuple<{ save:String, err:Error }> {
        try {
            return { save: Json.stringify(interpreter.save()), err: null };
        }
        catch (e:Any) {
            return { save: "", err: GoLibrary.toError(e) };
        }
    }

    @:go.Export
    public function restore(save:String):Error {
        try {
            interpreter.restore(Json.parse(save));
            return null;
        }
        catch (e:Any) {
            return GoLibrary.toError(e);
        }
    }

    @:go.Export
    public function stateField(name:String):Dynamic {
        return interpreter.getStateField(name);
    }

    @:go.Export
    public function setStateField(name:String, value:Dynamic):Void {
        interpreter.setStateField(name, value);
    }

    @:go.Export
    public function characterField(character:String, name:String):Dynamic {
        return interpreter.getCharacterField(character, name);
    }

    @:go.Export
    public function setCharacterField(character:String, name:String, value:Dynamic):Void {
        interpreter.setCharacterField(character, name, value);
    }

}
