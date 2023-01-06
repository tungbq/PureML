



def merge_step_with_value(values_dict, step):
    value_dict_with_step = {}

    for key, value in values_dict.items():
        value_dict_with_step[key] = {
                                        str(step) : value   #Check if this has to be replaced by ordered dict to preserve step order
                                    }
    
    return value_dict_with_step





def update_step_dict(values_dict_existing, values_dict_new):
    for key_new in values_dict_new.keys():
        if key_new in values_dict_existing.keys():

            steps_existing = list(values_dict_existing[key_new].keys())


            print(steps_existing, list(values_dict_new[key_new].keys()))

            step_new = list(values_dict_new[key_new].keys())[0]
            value_new = values_dict_new[key_new][step_new]

            if step_new in steps_existing:
                values_dict_existing[key_new][str(step_new)] = value_new
            else:
            
                values_dict_existing[key_new].update(
                                            {
                                                str(step_new) : value_new
                                            }
                                                )

        else:
            values_dict_existing[key_new].update(
                                        {
                                            str(step_new) : value_new
                                        }
                                            )


    return values_dict_existing
