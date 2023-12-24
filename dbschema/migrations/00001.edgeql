CREATE MIGRATION m1oh44setm2r7e236w5vctufmvrhrt7qv4csnkzc5mcffdaoljf7mq
    ONTO initial
{
  CREATE SCALAR TYPE default::NotificationType EXTENDING enum<FOLLOWER>;
  CREATE TYPE default::Notification {
      CREATE REQUIRED PROPERTY message: std::str;
      CREATE REQUIRED PROPERTY type: default::NotificationType;
  };
  CREATE TYPE default::User {
      CREATE MULTI LINK followers: default::User;
      CREATE MULTI LINK following: default::User;
      CREATE REQUIRED PROPERTY email: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE REQUIRED PROPERTY password: std::str;
  };
  CREATE TYPE default::Post {
      CREATE LINK author: default::User;
      CREATE REQUIRED PROPERTY content: std::str;
      CREATE REQUIRED PROPERTY title: std::str;
  };
};
