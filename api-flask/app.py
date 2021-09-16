from flask import Flask, jsonify, request
app = Flask(__name__)


@app.route('/')
def health_check():
    return jsonify({"status": "ok"})


@app.route('/tarea3-201801266', methods=['POST'])
def tarea3():
    return jsonify({"you_say": request.json})


if __name__ == '__main__':
    app.run(debug=True)
