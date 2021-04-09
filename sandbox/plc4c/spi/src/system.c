/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#include <plc4c/plc4c.h>
#include <plc4c/system.h>
#include <stdlib.h>
#include <string.h>

#include "plc4c/spi/system_private.h"
#include "plc4c/spi/types_private.h"

// As we're doing some operations where byte-order is important, we need this
// little helper to find out if we're on a big- or little-endian machine.
const int hurz = 1;
bool plc4c_is_bigendian() {
  return (((char) &hurz) == 0);
}

static void delete_driver_list_element(plc4c_list_element *driver_element) {
  plc4c_driver *driver = (plc4c_driver *)driver_element->value;
  // plc4c_driver_destroy(driver);
}

static void delete_transport_list_element(
    plc4c_list_element *transport_element) {
  plc4c_transport *transport = (plc4c_transport *)transport_element->value;
  // plc4c_transport_destroy(transport);
}
static void delete_connection_list_element(
    plc4c_list_element *connection_element) {
  plc4c_connection *connection = (plc4c_connection *)connection_element->value;
  plc4c_connection_destroy(connection);
  connection_element->value = NULL;
}
static void delete_task_list_element(plc4c_list_element *task_list_element) {
  plc4c_driver *driver = (plc4c_driver *)task_list_element->value;
  // plc4c_driver_destroy(driver);
}

plc4c_return_code plc4c_system_create(plc4c_system **system) {
  plc4c_system *new_system = malloc(sizeof(plc4c_system));

  plc4c_list *new_list = NULL;
  plc4c_utils_list_create(&new_list);
  new_system->driver_list = new_list;
  plc4c_utils_list_create(&new_list);
  new_system->transport_list = new_list;
  plc4c_utils_list_create(&new_list);
  new_system->connection_list = new_list;
  plc4c_utils_list_create(&new_list);
  new_system->task_list = new_list;

  new_system->on_driver_load_success_callback = NULL;
  new_system->on_driver_load_failure_callback = NULL;
  new_system->on_connect_success_callback = NULL;
  new_system->on_connect_failure_callback = NULL;
  new_system->on_disconnect_success_callback = NULL;
  new_system->on_disconnect_failure_callback = NULL;
  new_system->on_loop_failure_callback = NULL;

  *system = new_system;
  return OK;
}

void plc4c_system_destroy(plc4c_system *system) {
  // TODO: So some more cleaning up ...
  plc4c_utils_list_delete_elements(system->driver_list,
                                   &delete_driver_list_element);
  plc4c_utils_list_delete_elements(system->transport_list,
                                   &delete_transport_list_element);
  plc4c_utils_list_delete_elements(system->connection_list,
                                   &delete_connection_list_element);
  plc4c_utils_list_delete_elements(system->task_list,
                                   &delete_task_list_element);
  free(system->driver_list);
  free(system->transport_list);
  free(system->connection_list);
  free(system->task_list);
  free(system);
}

plc4c_list *plc4c_system_get_task_list(plc4c_system *system) {
  return system->task_list;
}

void plc4c_system_set_task_list(plc4c_system *system, plc4c_list *task_list) {
  system->task_list = task_list;
}

void plc4c_system_set_on_driver_load_success_callback(
    plc4c_system *system,
    plc4c_system_on_driver_load_success_callback callback) {
  system->on_driver_load_success_callback = callback;
}

void plc4c_system_set_on_driver_load_failure_callback(
    plc4c_system *system,
    plc4c_system_on_driver_load_failure_callback callback) {
  system->on_driver_load_failure_callback = callback;
}

void plc4c_system_set_on_connect_success_callback(
    plc4c_system *system, plc4c_system_on_connect_success_callback callback) {
  system->on_connect_success_callback = callback;
}

void plc4c_system_set_on_connect_failure_callback(
    plc4c_system *system, plc4c_system_on_connect_failure_callback callback) {
  system->on_connect_failure_callback = callback;
}

void plc4c_system_set_on_disconnect_success_callback(
    plc4c_system *system,
    plc4c_system_on_disconnect_success_callback callback) {
  system->on_disconnect_success_callback = callback;
}

void plc4c_system_set_on_disconnection_failure_callback(
    plc4c_system *system,
    plc4c_system_on_disconnect_failure_callback callback) {
  system->on_disconnect_failure_callback = callback;
}

void plc4c_system_set_on_loop_failure_callback(
    plc4c_system *system, plc4c_system_on_loop_failure_callback callback) {
  system->on_loop_failure_callback = callback;
}

plc4c_return_code plc4c_system_add_driver(plc4c_system *system,
                                          plc4c_driver *driver) {
  // If the system is not initialized, return an error.
  // There is nothing we can do here.
  if (system == NULL) {
    return INTERNAL_ERROR;
  }

  plc4c_utils_list_insert_tail_value(system->driver_list, driver);

  return OK;
}

plc4c_return_code plc4c_system_add_transport(plc4c_system *system,
                                             plc4c_transport *transport) {
  // If the system is not initialized, return an error.
  // There is nothing we can do here.
  if (system == NULL) {
    return INTERNAL_ERROR;
  }

  plc4c_utils_list_insert_tail_value(system->transport_list, transport);

  return OK;
}

plc4c_return_code plc4c_system_add_connection(plc4c_system *system,
                                              plc4c_connection *connection) {
  // If the system is not initialized, return an error.
  // There is nothing we can do here.
  if (system == NULL) {
    return INTERNAL_ERROR;
  }

  plc4c_utils_list_insert_tail_value(system->connection_list, connection);

  return OK;
}

void plc4c_system_remove_connection(plc4c_system *system,
                                    plc4c_connection *connection) {
  if (system == NULL || connection == NULL) {
    return;
  }
  plc4c_list_element *element = plc4c_utils_list_find_element_by_item(
      system->connection_list, connection);
  if (element != NULL) {
    plc4c_utils_list_remove(system->connection_list, element);
  }
  free(element);
}

plc4c_return_code plc4c_system_init(plc4c_system *system) {
  // Nothing to really do at the moment.

  return OK;
}

void plc4c_system_shutdown(plc4c_system *system) {}

plc4c_return_code plc4c_system_create_connection(
    char *connection_string, plc4c_connection **connection) {
  // Count the number of colons and question-marks so we know which pattern to
  // use for matching and how large the arrays for containing the different
  // segments should be.
  int num_colons = 0;
  int num_question_marks = 0;
  char *protocol_code = NULL;
  char *transport_code = NULL;
  char *transport_connect_information = NULL;
  char *parameters = NULL;
  int start_segment_index = 0;
  char *start_segment = connection_string;
  // The connection string has two parts ... the first, where colons are the
  // delimiters and the second where a question mark is the delimiter.
  enum mode {
    SEARCHING_FOR_COLONS,
    SEARCHING_FOR_QUESTION_MARKS
  } mode = SEARCHING_FOR_COLONS;
  for (int i = 0; i <= strlen(connection_string); i++) {
    // If we're in the first part of the connection string ... watch out for
    // colons.
    if (mode == SEARCHING_FOR_COLONS) {
      // If we encounter a colon, depending on the number of colons already
      // found, save the information in either the protocol code or transport
      // code variable.
      if (*(connection_string + i) == ':') {
        num_colons++;
        // The first colon delimits the protocol-code.
        if (num_colons == 1) {
          // Allocate enough memory to hold the sub-string.
          protocol_code =
              malloc(sizeof(char) * ((i - start_segment_index) + 1));
          *(protocol_code + (i - start_segment_index)) = '\0';
          // Copy the sub-string to the freshly allocated memory area.
          strncpy(protocol_code, start_segment, (i - start_segment_index));

          // Set the start of the next segment to directly after the colon.
          start_segment_index = i + 1;
          start_segment = connection_string + start_segment_index;

          // If the following character would be a slash, we're probably
          // finished and no transport code is provided. If this is the case,
          // ensure it's actually a double-slash and if this is the case switch
          // to the question-mark searching mode.
          if (*start_segment == '/') {
            if (*(start_segment + 1) == '/') {
              mode = SEARCHING_FOR_QUESTION_MARKS;
              start_segment_index += 2;
              start_segment += 2;
            } else {
              return INVALID_CONNECTION_STRING;
            }
          }
        }
        // If we encountered a second colon, this is the transport code.
        else if (num_colons == 2) {
          // Allocate enough memory to hold the sub-string.
          transport_code =
              malloc(sizeof(char) * ((i - start_segment_index) + 1));
          *(transport_code + (i - start_segment_index)) = '\0';
          // Copy the sub-string to the freshly allocated memory area.
          strncpy(transport_code, start_segment, (i - start_segment_index));

          // Set the start of the next segment to directly after the colon.
          start_segment_index = i + 1;
          start_segment = connection_string + start_segment_index;

          // The transport code is always followed by "://". So check if this is
          // the case. If it is, switch to question-mark searching mode.
          if ((*start_segment != '/') || (*(start_segment + 1) != '/')) {
            return INVALID_CONNECTION_STRING;
          }
          mode = SEARCHING_FOR_QUESTION_MARKS;

          // Bump the start of the segment to after the double slashes.
          start_segment_index += 2;
          start_segment += 2;
        } else {
          return INVALID_CONNECTION_STRING;
        }
      }
    }
    // If we're in the second part, look for question marks.
    else {
      // The question-mark separates the transport connect information from the
      // parameters.
      if (*(connection_string + i) == '?') {
        num_question_marks++;
        // Only one question-mark is allowed in connection strings.
        if (num_question_marks > 1) {
          return INVALID_CONNECTION_STRING;
        }

        // Allocate enough memory to hold the sub-string.
        transport_connect_information =
            malloc(sizeof(char) * ((i - start_segment_index) + 1));
        *(transport_connect_information + (i - start_segment_index)) = '\0';
        // Copy the sub-string to the freshly allocated memory area.
        strncpy(transport_connect_information, start_segment,
                (i - start_segment_index));

        // Set the start of the next segment to directly after the
        // question-mark.
        start_segment_index = i + 1;
        start_segment = connection_string + start_segment_index;
      }
      // This is the last character ... finish up the last loose end.
      if (i == strlen(connection_string)) {
        // If no question-mark has been encountered, this connection string
        // doesn't have one and the remaining part is simply the transport
        // connect information.
        if (num_question_marks == 0) {
          transport_connect_information =
              malloc(sizeof(char) * ((i - start_segment_index)) + 2);
          *(transport_connect_information + (i - start_segment_index) + 1) = '\0';
          strncpy(transport_connect_information, start_segment,
                  (i - start_segment_index) + 1);
        }
        // I a question-mark was found, this is the parameters section.
        else {
          parameters = malloc(sizeof(char) * (i - start_segment_index) + 2);
          *(parameters + (i - start_segment_index) + 1) = '\0';
          strncpy(parameters, start_segment, (i - start_segment_index) + 1);
        }
      }
    }
  }
  if (num_colons == 0) {
    return INVALID_CONNECTION_STRING;
  }

  // Initialize a new connection data-structure with the parsed information.
  plc4c_connection *new_connection = malloc(sizeof(plc4c_connection));
  plc4c_connection_initialize(new_connection);
  plc4c_connection_set_connection_string(new_connection, connection_string);
  plc4c_connection_set_protocol_code(new_connection, protocol_code);
  plc4c_connection_set_transport_code(new_connection, transport_code);
  plc4c_connection_set_transport_connect_information(
      new_connection, transport_connect_information);
  plc4c_connection_set_parameters(new_connection, parameters);

  *connection = new_connection;
  return OK;
}

plc4c_return_code plc4c_system_connect(plc4c_system *system,
                                       char *connection_string,
                                       plc4c_connection **connection) {
  // Parse the connection string and initialize some of the connection field
  // variables from this.
  plc4c_connection *new_connection = NULL;
  plc4c_return_code result =
      plc4c_system_create_connection(connection_string, &new_connection);
  if (result != OK) {
    return result;
  }

  plc4c_connection_set_system(new_connection, system);

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // Find a matching driver from the driver-list
  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

  // If no driver is available at all this is definitely a developer error,
  // so we output a special error code for this case
  if (plc4c_utils_list_empty(system->driver_list)) {
    return NO_DRIVER_AVAILABLE;
  }
  plc4c_list_element *cur_driver_list_element = system->driver_list->tail;
  do {
    plc4c_driver *cur_driver = (plc4c_driver *)cur_driver_list_element->value;
    if (strcmp(cur_driver->protocol_code,
               plc4c_connection_get_protocol_code(new_connection)) == 0) {
      // Set the driver reference in the new connection.
      plc4c_connection_set_driver(new_connection, cur_driver);

      // If no transport was selected, use the drivers default transport (if it
      // exists).
      if (plc4c_connection_get_transport_code(new_connection) == NULL) {
        if (cur_driver->default_transport_code != NULL) {
          plc4c_connection_set_transport_code(
              new_connection, cur_driver->default_transport_code);
        }
      }
      break;
    }
    cur_driver_list_element = cur_driver_list_element->next;
  } while (cur_driver_list_element != NULL);

  // If the driver property is still NULL, the desired driver was not found.
  if (plc4c_connection_get_driver(new_connection) == NULL) {
    return UNKNOWN_DRIVER;
  }

  // Return an error if the user didn't specify a transport and the driver
  // doesn't have a default one.
  if (plc4c_connection_get_transport_code(new_connection) == NULL) {
    return UNSPECIFIED_TRANSPORT;
  }

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // Find a matching transport from the transport-list
  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

  // If no transport is available at all this is definitely a developer error,
  // so we output a special error code for this case
  if (plc4c_utils_list_empty(system->transport_list)) {
    return NO_TRANSPORT_AVAILABLE;
  }
  plc4c_list_element *cur_transport_list_element = system->transport_list->tail;
  do {
    plc4c_transport *cur_transport =
        (plc4c_transport *)cur_transport_list_element->value;
    if (strcmp(cur_transport->transport_code,
               plc4c_connection_get_transport_code(new_connection)) == 0) {
      // Set the transport reference in the new connection.
      plc4c_connection_set_transport(new_connection, cur_transport);
      break;
    }
    cur_transport_list_element = cur_transport_list_element->next;
  } while (cur_transport_list_element != NULL);

  // If the transport property is still NULL, the desired transport was not
  // found.
  if (plc4c_connection_get_transport(new_connection) == NULL) {
    return UNKNOWN_TRANSPORT;
  }

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // Initialize a new connection task and schedule that.
  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

  plc4c_driver* driver = plc4c_connection_get_driver(new_connection);

  // Configure the driver configuration first.
  void* configuration = NULL;
  // TODO: Pass in the configuration options ...
  driver->configure_function(NULL, &configuration);
  plc4c_connection_set_configuration(new_connection, configuration);

  // TODO: Somehow let the driver inject default values which the transport can then pickup ...

  // Prepare a configuration data structure for the current transport.
  new_connection->transport->configure(new_connection->transport_connect_information, NULL, &new_connection->transport_configuration);

  // Create a new connection task.
  plc4c_system_task *new_connection_task = NULL;
  result = driver->connect_function(new_connection, &new_connection_task);
  if (result != OK) {
    return -1;
  }
  // Increment the number of running tasks for this connection.
  plc4c_connection_task_added(new_connection);
  plc4c_utils_list_insert_tail_value(system->task_list, new_connection_task);

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // Add the new connection to the systems connection-list.
  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

  result = plc4c_system_add_connection(system, new_connection);
  if (result != OK) {
    return result;
  }

  // Pass the new connection back.
  *connection = new_connection;

  return OK;
}

plc4c_return_code plc4c_system_loop(plc4c_system *system) {
  // If the task-queue is empty, just return.
  if (plc4c_utils_list_empty(system->task_list)) {
    return OK;
  }

  plc4c_list_element *cur_task_element =
      plc4c_utils_list_head(system->task_list);
  do {
    // Get the current element's system task.
    plc4c_system_task *cur_task = cur_task_element->value;

    // If the task is already completed, no need to do anything.
    if ((!cur_task->completed) && (cur_task->state_machine_function != NULL)) {
      // Pass the task itself to the state-machine function of this task.
      cur_task->state_machine_function(cur_task);
    }

    // If the current task is completed at the end, remove it from the
    // task_queue.
    if (cur_task->completed) {
      plc4c_utils_list_remove(system->task_list, cur_task_element);
      if (cur_task->connection != NULL) {
        plc4c_connection_task_removed(cur_task->connection);
      }
      // TODO: clean up the memory of the cur_task_element and cur_task
    }

    cur_task_element = cur_task_element->next;
  } while (cur_task_element != NULL);

  return OK;
}

char* list_to_string(plc4c_list* list) {
  uint8_t string_length = plc4c_utils_list_size(list);
  char* chars = malloc(sizeof(char) * (string_length + 1));
  if(chars == NULL) {
    return NULL;
  }
  char* cur_pos = chars;
  plc4c_list_element* cur_element = list->tail;
  while(cur_element != NULL) {
    char cur_char = *((char*)(cur_element->value));
    *cur_pos = cur_char;
    cur_element = cur_element->next;
    cur_pos++;
  }
  // Terminate the string.
  *cur_pos = '\0';
  return chars;
}

