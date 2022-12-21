import os
import sys
import face_recognition
import numpy as np


def register(face_path: str) -> str:
    face_image = face_recognition.load_image_file(face_path)
    face = face_recognition.face_encodings(face_image)[0]

    filename = os.getcwd() + "/cv_service/service/face_verification/face_descriptor.txt"
    np.savetxt(filename, face)
    return filename


def authorize(unknown_path: str, known_path: str) -> str:
    unknown_image = face_recognition.load_image_file(unknown_path)
    unknown_face = face_recognition.face_encodings(unknown_image)[0]

    known_face = np.loadtxt(known_path)

    compare_result = face_recognition.compare_faces([known_face], unknown_face)
    if compare_result[0]:
        return "true"
    return "false"


def main(argv):
    option = argv[1]
    result = ""
    if option == "--register":
        result = register(argv[2])
    elif option == "--authorize":
        result = authorize(argv[2], argv[3])
    print(result)


if __name__ == "__main__":
    main(sys.argv)
