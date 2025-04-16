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
        final_result[el[1]] = faculty_ids[str(el[0])]

    with open(RESULT_FILE, 'w') as f:
        json.dump(final_result, f)

    fp_index, text_result = 0, sorted(final_result.items())
    for i, val in enumerate(text_result):
        if val[0] > 0:
            fp_index = i
            break;

    text_result = text_result[fp_index:] + text_result[:fp_index]
    with open(TEXT_RESULT_FILE, 'w') as f:
        for key, val in text_result:
            f.write(str(key) + " : " + val + "\n")
