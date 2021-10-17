from flask import Flask, jsonify, request
from rekognition import detect_labels

app = Flask(__name__)


@app.route('/')
def health_check():
    return jsonify({"status": "ok"})


@app.route('/detect-labels', methods=['POST'])
def tarea3():
    return jsonify(detect_labels(request.json['image']))


if __name__ == '__main__':
    app.run(debug=True)
