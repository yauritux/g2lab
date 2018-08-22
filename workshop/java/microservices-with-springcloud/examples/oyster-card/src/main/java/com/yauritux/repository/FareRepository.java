package com.yauritux.repository;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Fare;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface FareRepository extends JpaRepository<Fare, Long> {

	List<Fare> findByJourneyType(JourneyType journeyType);
}
