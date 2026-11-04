package loreline.lib;

import haxe.Json;
import loreline.Interpreter;
import loreline.Loreline;
import loreline.Script;

@:keep
class GoLibrary {

    public static var error:String = null;

    static function main() {}

    static function fail(e:Any):Void {
        error = Std.string(e);
    }

    static function beatOrNull(beat:String):String {
        return beat == null || beat.length == 0 ? null : beat;
    }

    static function dialogueBridge(dialogue:(Interpreter, String, String, Array<TextTag>, () -> Void) -> Void):DialogueHandler {
        return (interpreter, character, text, tags, callback) -> dialogue(interpreter, character ?? "", text, tags, callback);
    }

    public static function parse(source:String, filePath:String, readFile:String -> String):Script {
        error = null;
        try {
            final script = filePath == null || filePath.length == 0
                ? Loreline.parse(source)
                : Loreline.parse(source, filePath, (path, callback) -> callback(readFile(path)));
            if (script == null) {
                final e = Loreline.lastError();
                error = e != null ? Std.string(e) : "Failed to parse script";
            }
            return script;
        }
        catch (e:Any) {
            fail(e);
            return null;
        }
    }

    public static function play(
        script:Script,
        dialogue:(Interpreter, String, String, Array<TextTag>, () -> Void) -> Void,
        choice:(Interpreter, Array<ChoiceOption>, Int -> Void) -> Void,
        finish:Interpreter -> Void,
        beat:String
    ):Interpreter {
        error = null;
        try {
            return Loreline.play(script, dialogueBridge(dialogue), choice, finish, beatOrNull(beat));
        }
        catch (e:Any) {
            fail(e);
            return null;
        }
    }

    public static function resume(
        script:Script,
        dialogue:(Interpreter, String, String, Array<TextTag>, () -> Void) -> Void,
        choice:(Interpreter, Array<ChoiceOption>, Int -> Void) -> Void,
        finish:Interpreter -> Void,
        saveJson:String,
        beat:String
    ):Interpreter {
        error = null;
        try {
            return Loreline.resume(script, dialogueBridge(dialogue), choice, finish, Json.parse(saveJson), beatOrNull(beat));
        }
        catch (e:Any) {
            fail(e);
            return null;
        }
    }

    public static function start(interpreter:Interpreter, beat:String):Void {
        error = null;
        try {
            interpreter.start(beatOrNull(beat));
        }
        catch (e:Any) {
            fail(e);
        }
    }

    public static function save(interpreter:Interpreter):String {
        error = null;
        try {
            return Json.stringify(interpreter.save());
        }
        catch (e:Any) {
            fail(e);
            return null;
        }
    }

    public static function restore(interpreter:Interpreter, saveJson:String):Void {
        error = null;
        try {
            interpreter.restore(Json.parse(saveJson));
        }
        catch (e:Any) {
            fail(e);
        }
    }

    public static function getStateField(interpreter:Interpreter, name:String):Dynamic {
        return interpreter.getStateField(name);
    }

    public static function setStateField(interpreter:Interpreter, name:String, value:Dynamic):Void {
        interpreter.setStateField(name, value);
    }

    public static function getCharacterField(interpreter:Interpreter, character:String, name:String):Dynamic {
        return interpreter.getCharacterField(character, name);
    }

    public static function setCharacterField(interpreter:Interpreter, character:String, name:String, value:Dynamic):Void {
        interpreter.setCharacterField(character, name, value);
    }

    public static function print(script:Script):String {
        return Loreline.print(script);
    }

}
