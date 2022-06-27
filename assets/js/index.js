// retrieve elements
const runCodeBtn = document.querySelector(".editor__run");
const resetCodeBtn = document.querySelector(".editor__reset");

var codeEditor = ace.edit("editorCode");

var editorLib = {
  init() {
    codeEditor.setTheme("ace/theme/cloud9_night_low_color");
    codeEditor.session.setMode("ace/mode/golang");
    codeEditor.insert("//Write your code here")
  },
};

runCodeBtn.addEventListener("click", () => {
  let code = codeEditor.getValue();
  let data = {
    "script": code,
    "language": "go",
    "versionIndex": "0"
  }
  hitAPI(data);
});

resetCodeBtn.addEventListener("click", () => {
  codeEditor.setValue("");
});


function strReplace(output){
  var myStr = output;
  var newStr = myStr.replace(/\n/g, "<br />");
  return newStr
}

async function hitAPI(data) {
  try {
    const response = await fetch('http://localhost:8080/v1/execute', {
      method: 'POST',
      headers: {
        accept: 'application/json',
      },
      body: JSON.stringify(data)
    });

    if (!response.ok) {
      throw new Error(`Error! status: ${response.status}`);
    }

    var result = await response.json();
    console.log(result)

    document.getElementById("demo").innerHTML = strReplace(result.output);
  } catch (err) {
    console.log(err);
  }
}

editorLib.init();

