package loreline.lib;

import go.Error;
import go.Tuple;
import haxe.Json;
import loreline.Interpreter;
import loreline.Loreline;
import loreline.Script;

typedef GoDialogueHandler = (GoInterpreter, String, String, Array<TextTag>, () -> Void) -> Void;

typedef GoChoiceHandler = (GoInterpreter, Array<ChoiceOption>, Int -> Void) -> Void;

typedef GoFinishHandler = GoInterpreter -> Void;

class GoScript {

    final script:Script;

    public function new(script:Script) {
        this.script = script;
    }

    @:go.Export
    @:go.Tuple('interpreter', 'err')
    public function play(dialogue:GoDialogueHandler, choice:GoChoiceHandler, finish:GoFinishHandler, beat:String):Tuple<{ interpreter:GoInterpreter, err:Error }> {
        final wrapper = new GoInterpreter();
        try {
            wrapper.interpreter = Loreline.play(script, dialogueBridge(wrapper, dialogue), choiceBridge(wrapper, choice), finishBridge(wrapper, finish), GoInterpreter.beatOrNull(beat));
            return { interpreter: wrapper, err: null };
        }
        catch (e:Any) {
            return { interpreter: null, err: GoLibrary.toError(e) };
        }
    }

    @:go.Export
    @:go.Tuple('interpreter', 'err')
    public function resume(dialogue:GoDialogueHandler, choice:GoChoiceHandler, finish:GoFinishHandler, save:String, beat:String):Tuple<{ interpreter:GoInterpreter, err:Error }> {
        final wrapper = new GoInterpreter();
        try {
            wrapper.interpreter = Loreline.resume(script, dialogueBridge(wrapper, dialogue), choiceBridge(wrapper, choice), finishBridge(wrapper, finish), Json.parse(save), GoInterpreter.beatOrNull(beat));
            return { interpreter: wrapper, err: null };
        }
        catch (e:Any) {
            return { interpreter: null, err: GoLibrary.toError(e) };
        }
    }

    @:go.Export
    public function format():String {
        return Loreline.print(script);
    }

    static function dialogueBridge(wrapper:GoInterpreter, dialogue:GoDialogueHandler):DialogueHandler {
        return (interpreter, character, text, tags, callback) -> {
            wrapper.interpreter = interpreter;
            dialogue(wrapper, character ?? "", text, tags, callback);
        };
    }

    static function choiceBridge(wrapper:GoInterpreter, choice:GoChoiceHandler):ChoiceHandler {
        return (interpreter, options, callback) -> {
            wrapper.interpreter = interpreter;
            choice(wrapper, options, callback);
        };
    }

    static function finishBridge(wrapper:GoInterpreter, finish:GoFinishHandler):FinishHandler {
        return interpreter -> {
            wrapper.interpreter = interpreter;
            finish(wrapper);
        };
    }

}
