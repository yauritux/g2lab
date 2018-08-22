package com.yauritux.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.yauritux.model.entity.Station;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface StationRepository extends JpaRepository<Station, Long> {

	Station findByName(String name);
}
