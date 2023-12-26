CREATE MIGRATION m1w24s3ath5czb5pnjzp2h47xd66vf36hw3ygodlljxitxrugg4d5a
    ONTO m1oh44setm2r7e236w5vctufmvrhrt7qv4csnkzc5mcffdaoljf7mq
{
  ALTER TYPE default::Notification {
      CREATE LINK user: default::User;
  };
};
