INSERT INTO `tasks` (`title`, `description`) VALUES ("sample-task-01", "this is a description of sample-task-01");
INSERT INTO `tasks` (`title`) VALUES ("sample-task-02");
INSERT INTO `tasks` (`title`, `deadline`, `is_done`) VALUES ("sample-task-03", "2022-01-01 00:00:00", true);
INSERT INTO `tasks` (`title`, `deadline`) VALUES ("sample-task-04", "2021-12-31 23:59:00");
INSERT INTO `tasks` (`title`, `deadline`) VALUES ("final-report", "2022-11-29 23:59:00");
INSERT INTO `tasks` (`title`, `is_done`) VALUES ("sample-task-06", true);
INSERT INTO `tasks` (`title`, `deadline`, `is_done`) VALUES ("sample-task-07", "2023-01-01 00:00:00", true);

INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 1);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 2);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 3);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 4);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 5);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 6);
INSERT INTO `ownership` (`user_id`, `task_id`) VALUES (1, 7);