import json

INPUT_FILE       = "preferences_clean.json"
OUTPUT_FILE      = "matching_output.json"
RESULT_FILE      = "result.json"
TEXT_RESULT_FILE = "result_text.txt"

def get_faculty_id_map(json_meta_data : dict):
    return json_meta_data["user"]

def get_matching_formatted(json_matching_output: dict):
    res = []
    for matching in json_matching_output["Matchings"]:
        res.append((matching["user_id"], matching["course_sem_id"]))
    return res

if __name__ == "__main__":
    json_meta_data = None
    json_matching_output = None
    with open(INPUT_FILE, 'r') as f:
        json_meta_data = json.load(f)
    with open(OUTPUT_FILE, 'r') as f:
        json_matching_output = json.load(f)

    faculty_ids = get_faculty_id_map(json_meta_data)
    res         = get_matching_formatted(json_matching_output)

    final_result = {}
    for el in res:
        final_result[faculty_ids[str(el[0])]] = el[1]

    with open(RESULT_FILE, 'w') as f:
        json.dump(final_result, f)

    with open(TEXT_RESULT_FILE, 'w') as f:
        for key, val in final_result.items():
            f.write(key + " : " + str(val) + "\n")
